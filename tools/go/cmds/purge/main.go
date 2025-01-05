package purge

import (
	"context"
	"fmt"
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/config"
	"jgttech/dotfiles/src/exec"
	"os"

	"github.com/urfave/cli/v3"
)

func Command(build *config.BuildJson) *cli.Command {
	cfg := build.Config
	home := os.Getenv("HOME")

	return &cli.Command{
		Name:  "purge",
		Usage: "Unlinks all the stow links and CLI, then deletes the source code from the system.",
		Action: func(ctx context.Context, c *cli.Command) error {
			// Detect and unlink all the packages.
			cmd := exec.Cmd(fmt.Sprintf("%s uninstall", cfg.Binary))

			if err := cmd.Run(); err != nil {
				return err
			}

			// Unlink the build.
			cmd = exec.Cmd(fmt.Sprintf("stow -t %s -D .build", home))
			cmd.Dir = cfg.Base

			assert.Will(cmd.Run())
			return nil
		},
	}
}
