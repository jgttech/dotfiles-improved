package env

import (
	"context"
	"fmt"
	"jgttech/dotfiles/src/config"

	"github.com/urfave/cli/v3"
)

func Command(build *config.BuildJson) *cli.Command {
	cfg := build.Config

	return &cli.Command{
		Name:  "env",
		Usage: "",
		Action: func(ctx context.Context, c *cli.Command) error {
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
