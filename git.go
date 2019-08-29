package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type (
	// Git config
	Config struct {
		Enable bool   // Git clone enable
		Url    string // Git repo url
		Out    string // Export package
		Token  string // Gitlab Personal Access Token
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
)

// Exec executes the plugin step
func (p Plugin) Exec() error {

	// git clone configuration
	if p.Config.Enable {
		cmd := commandClone(p.Config)
		trace(cmd)
		out, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", out)
	} else {
		fmt.Println("enable = false,Ignore pull configuration")
	}

	// git check and write packages file
	if p.Check.Enable {
		cmd := commandCheckFileList(p.Check)
		trace(cmd)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		var pkglist []string
		files := strings.Split(string(out), "\n")
		for _, file := range files {
			pkg := strings.Split(file, "/")[0]
			pkglist = append(pkglist, pkg)
		}
		recordFiles(removeDuplicateElement(pkglist))
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
		fmt.Println("no 'git' program on path")
	}
	return gitProgram
}

// commandClone git clone configuration
func commandClone(config Config) *exec.Cmd {
	url := strings.Replace(config.Url, "https://", "", 1)
	clone_url := fmt.Sprintf("https://oauth2:%s@%s", config.Token, url)
	return exec.Command(
		commandGit(),
		"clone",
		clone_url,
		config.Out,
	)
}

// commandCheckFileList get diff files list command
func commandCheckFileList(check Check) *exec.Cmd {
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
func recordFiles(pkglist []string) string {
	target := strings.Join(pkglist, ",")
	fmt.Fprintf(os.Stdout, "+ %s\n", target)
	content := []byte(target)
	err := ioutil.WriteFile("git.txt", content, 0666)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
		os.Exit(0)
	}
	return fmt.Sprintf("deploy package: %s", target)
}

// trace writes each command to stdout with the command wrapped in an xml
// tag so that it can be extracted and displayed in the logs.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}
