package purge

import (
	"context"
	"fmt"
	"jgttech/dotfiles/src/config"
	// "jgttech/dotfiles/src/exec"

	"github.com/urfave/cli/v3"
)

func Command(build *config.BuildJson) *cli.Command {
	cfg := build.Config

	return &cli.Command{
		Name:  "purge",
		Usage: "",
		Action: func(ctx context.Context, c *cli.Command) error {
			// exec.Cmd("")
			fmt.Printf("%#v\n", cfg)

			return nil
		},
	}
}
