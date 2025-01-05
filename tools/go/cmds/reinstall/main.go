package reinstall

import (
	"context"
	"fmt"
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/config"
	"jgttech/dotfiles/src/exec"
	// "os"

	"github.com/urfave/cli/v3"
)

func Command(build *config.BuildJson) *cli.Command {
	cfg := build.Config
	// home := os.Getenv("HOME")

	return &cli.Command{
		Name:  "reinstall",
		Usage: "Uninstall the packages, rebuild the CLI, and re-install the packages.",
		Action: func(ctx context.Context, c *cli.Command) error {
			cmd := exec.Cmd(fmt.Sprintf("%s uninstall", cfg.Binary))

			// Uninstall the packages.
			assert.Will(cmd.Run())

			// cmd = exec.Cmd(fmt.Sprintf("stow -t %s .build", home))
			// cmd.Dir = cfg.Base

			// Uninstall the CLI.
			// assert.Will(cmd.Run())

			// Delete the build.
			// assert.Will(os.RemoveAll(cfg.Build))

			return nil
		},
	}
}
