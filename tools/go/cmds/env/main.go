package env

import (
	ctx "context"
	"fmt"
	"jgttech/dotfiles/src/context"

	"github.com/urfave/cli/v3"
)

func Command(etx *context.ExecutionContext) *cli.Command {
	cfg := etx.Build.Config

	return &cli.Command{
		Name:  "env",
		Usage: "",
		Action: func(ctx ctx.Context, c *cli.Command) error {
			fmt.Println("base.....:", cfg.Base)
			fmt.Println("build....:", cfg.Build)
			fmt.Println("tools....:", cfg.Tools)
			fmt.Println("binary...:", cfg.Binary)
			fmt.Println("symlink..:", cfg.Symlink)
			fmt.Println("language.:", cfg.Language)

			return nil
		},
	}
}
