package rebuild

import (
	"context"
	"fmt"
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/config"
	"jgttech/dotfiles/src/exec"
	"os"
	"path"

	"github.com/urfave/cli/v3"
)

func Command(build *config.BuildJson) *cli.Command {
	cfg := build.Config
	home := os.Getenv("HOME")

	return &cli.Command{
		Name:  "rebuild",
		Usage: "Uninstall the packages, rebuild the CLI, and re-install the packages.",
		Action: func(ctx context.Context, c *cli.Command) error {
			cmd := exec.Cmd(fmt.Sprintf("%s uninstall", cfg.Binary), exec.Stdio)

			// Uninstall the packages.
			assert.Will(cmd.Run())

			cmd = exec.Cmd(fmt.Sprintf("stow -t %s -D .build", home), exec.Stdio)
			cmd.Dir = cfg.Base

			// Uninstall the CLI.
			assert.Will(cmd.Run())

			// Delete the build.
			assert.Will(os.RemoveAll(path.Join(
				cfg.Build,
				cfg.Symlink,
				cfg.Binary,
			)))

			cmd = exec.Cmd(
				fmt.Sprintf(
					"bun --cwd %s build %s %s",
					cfg.Tools,
					cfg.Symlink,
					cfg.Binary,
				),
				exec.Stdio,
			)
			cmd.Dir = cfg.Base

			// Rebuild the CLI.
			assert.Will(cmd.Run())

			cmd = exec.Cmd(fmt.Sprintf("stow -t %s .build", home), exec.Stdio)
			cmd.Dir = cfg.Base

			// Install the CLI.
			assert.Will(cmd.Run())

			cmd = exec.Cmd(fmt.Sprintf("%s install", cfg.Binary), exec.Stdio)
			cmd.Dir = home

			// Run the install command.
			assert.Will(cmd.Run())

			return nil
		},
	}
}
