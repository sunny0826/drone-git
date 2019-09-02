package git

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type (
	// Git config
	Config struct {
		Url   string // Git repo url
		Out   string // Export package
		Token string // Gitlab Personal Access Token
	}
	// Check package
	Check struct {
		Enable bool     // Git check enable
		List   []string // check package list
		Commit string   // providers the commit sha for the current build
	}

	// Plugin defines the Docker plugin parameters.
	Plugin struct {
		Config Config // Git clone configuration
		Check  Check  // Git check configuration
	}
	Envfile struct {
		ConfigPkg string   `yaml:"configPkg"`
		CheckList []string `yaml:"checkList"`
	}
)

// Exec executes the plugin step
func (p Plugin) Exec() error {

	// git clone configuration
	envyaml := Envfile{}
	cloneCmd := envyaml.commandClone(p.Config)
	//trace(cmd)
	err := cloneCmd.Run()
	if err != nil {
		return fmt.Errorf("+ %s", err)
	}

	// git check and write packages file
	if p.Check.Enable {
		cmd := commandCheckFileList(p.Check)
		//trace(cmd)
		out, err := cmd.Output()
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
		}
		var pkglist []string
		files := strings.Split(string(out), "\n")
		for _, file := range files {
			pkg := strings.Split(file, "/")[0]
			if pkg != "" && len(strings.Split(pkg, ".")) == 1 {
				pkglist = append(pkglist, pkg)
			}
		}
		envyaml.recordFiles(removeDuplicateElement(pkglist), p.Config.Out)
	}

	return nil
}

func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// commandGit git command bin path
func commandGit() string {
	gitProgram, err := exec.LookPath("git")
	if err != nil {
		fmt.Fprintln(os.Stdout, "no 'git' program on path")
	}
	return gitProgram
}

// commandClone git clone configuration
func (env *Envfile) commandClone(config Config) *exec.Cmd {
	fmt.Fprintf(os.Stdout, "+ clone %s to %s\n", config.Url, config.Out)
	url := strings.Replace(config.Url, "https://", "", 1)
	clone_url := fmt.Sprintf("https://oauth2:%s@%s", config.Token, url)
	env.ConfigPkg = config.Out
	env.WriteYaml()
	return exec.Command(
		commandGit(),
		"clone",
		clone_url,
		config.Out,
	)
}

// commandCheckFileList get diff files list command
func commandCheckFileList(check Check) *exec.Cmd {
	fmt.Fprintf(os.Stdout, "+ check commit: %s\n", check.Commit)
	return exec.Command(
		commandGit(),
		"diff-tree",
		"--no-commit-id",
		"--name-only",
		"-r",
		check.Commit,
	)
}

// write diff list of commit
func (env *Envfile) recordFiles(pkglist []string, out string) {
	target := strings.Join(pkglist, ",")
	if len(pkglist) == 0 {
		fmt.Fprintln(os.Stdout, "+ no change packages \n")
	} else {
		fmt.Fprintf(os.Stdout, "+ change packages: %s\n", target)
	}
	//content := []byte(target)
	//env.ReadYaml("./env.yaml")
	//env.ConfigPkg = out
	env.CheckList = pkglist
	env.WriteYaml()

	//err := ioutil.WriteFile("git.txt", content, 0666)
	//if err != nil {
	//	fmt.Println("ioutil WriteFile error: ", err)
	//	os.Exit(0)
	//}
}

//func (c *Envfile) ReadYaml(f string) {
//	buffer, err := ioutil.ReadFile(f)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//	err = yaml.Unmarshal(buffer, &c)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//}

func (c *Envfile) WriteYaml() {
	buffer, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = ioutil.WriteFile("./env.yaml", buffer, 0777)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// trace writes each command to stdout with the command wrapped in an xml
// tag so that it can be extracted and displayed in the logs.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ run: %s\n", strings.Join(cmd.Args, " "))
}
