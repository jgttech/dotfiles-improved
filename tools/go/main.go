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
	cfg := config.NewBuildJson()
	app := &cli.Command{
		Name:  "dotfiles",
		Usage: "My personal dotfiles CLI.",
		Commands: []*cli.Command{
			env.Command(),
			install.Command(cfg),
			purge.Command(),
			uninstall.Command(cfg),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
