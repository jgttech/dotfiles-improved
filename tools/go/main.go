package main

import (
	"context"
	"os"

	"jgttech/dotfiles/cmds/env"
	"jgttech/dotfiles/cmds/install"
	"jgttech/dotfiles/cmds/purge"
	"jgttech/dotfiles/cmds/uninstall"
	"jgttech/dotfiles/src/config"

	"github.com/urfave/cli/v3"
)

func main() {
	build := config.NewBuildJson()
	app := &cli.Command{
		Name:  "dotfiles",
		Usage: "My personal dotfiles CLI.",
		Commands: []*cli.Command{
			env.Command(build),
			install.Command(build),
			purge.Command(build),
			uninstall.Command(build),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
