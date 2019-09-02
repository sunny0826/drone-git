package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sunny0826/drone-git/git"
	"github.com/urfave/cli"
	"os"
)

var (
	version = "unknown"
)

func main() {
	// Load env-file if it exists first
	if env := os.Getenv("PLUGIN_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := cli.NewApp()
	app.Name = "git plugin"
	app.Usage = "git plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "git-config.url",
			Usage:  "git config url",
			EnvVar: "PLUGIN_CONFIG_PATH",
		},
		cli.StringFlag{
			Name:   "access.tokens",
			Usage:  "Personal Access Token",
			EnvVar: "PLUGIN_TOKEN",
		},
		cli.StringFlag{
			Name:   "git-config.out",
			Usage:  "git config out",
			EnvVar: "PLUGIN_OUT",
		},
		cli.BoolFlag{
			Name:   "check.enable",
			Usage:  "git check enable",
			EnvVar: "PLUGIN_CHECK_ENABLE",
		},
		cli.StringSliceFlag{
			Name:   "check.list",
			Usage:  "git check list",
			EnvVar: "PLUGIN_CHECK_LIST",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "providers the commit sha for the current build",
			EnvVar: "DRONE_COMMIT_SHA",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := git.Plugin{
		Config: git.Config{
			Url:    c.String("git-config.url"),
			Out:    c.String("git-config.out"),
			Token:  c.String("access.tokens"),
		},
		Check: git.Check{
			Enable: c.Bool("check.enable"),
			Commit: c.String("commit.sha"),
			List:   c.StringSlice("check.list"),
		},
	}

	return plugin.Exec()
}
