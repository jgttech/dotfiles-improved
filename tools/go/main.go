package main

import (
	ctx "context"
	"fmt"
	"os"

	"jgttech/dotfiles/cmds/env"
	"jgttech/dotfiles/cmds/install"
	"jgttech/dotfiles/cmds/purge"
	"jgttech/dotfiles/cmds/rebuild"
	"jgttech/dotfiles/cmds/uninstall"
	"jgttech/dotfiles/src/context"

	"github.com/urfave/cli/v3"
)

func main() {
	etx := context.NewExecutionContext()
	fmt.Println("packages.:", etx.Packages())
	fmt.Println("tools....:", etx.Tools())

	app := &cli.Command{
		Name:  "dotfiles",
		Usage: "My personal dotfiles CLI.",
		Commands: []*cli.Command{
			env.Command(etx),
			install.Command(etx),
			purge.Command(etx),
			rebuild.Command(etx),
			uninstall.Command(etx),
		},
	}

	if err := app.Run(ctx.Background(), os.Args); err != nil {
		panic(err)
	}
}
