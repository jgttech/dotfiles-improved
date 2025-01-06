package main

import (
	ctx "context"
	"os"

	"jgttech/dotfiles/cmds/env"
	"jgttech/dotfiles/cmds/install"
	"jgttech/dotfiles/cmds/packages"
	"jgttech/dotfiles/cmds/purge"
	"jgttech/dotfiles/cmds/rebuild"
	"jgttech/dotfiles/cmds/reinstall"
	"jgttech/dotfiles/cmds/tools"
	"jgttech/dotfiles/cmds/uninstall"
	"jgttech/dotfiles/src/context"

	"github.com/urfave/cli/v3"
)

func main() {
	etx := context.NewExecutionContext()
	app := &cli.Command{
		Name:  "dotfiles",
		Usage: "My personal dotfiles CLI.",
		Commands: []*cli.Command{
			env.Command(etx),
			install.Command(etx),
			packages.Command(etx),
			purge.Command(etx),
			rebuild.Command(etx),
			reinstall.Command(),
			tools.Command(etx),
			uninstall.Command(etx),
		},
	}

	if err := app.Run(ctx.Background(), os.Args); err != nil {
		panic(err)
	}
}
