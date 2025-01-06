package purge

import (
	ctx "context"
	"fmt"
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/context"
	"jgttech/dotfiles/src/exec"
	"os"

	"github.com/urfave/cli/v3"
)

func Command(etx *context.ExecutionContext) *cli.Command {
	cfg := etx.Build.Config
	home := os.Getenv("HOME")

	return &cli.Command{
		Name:  "purge",
		Usage: "Uninstalls the packages and the CLI, as well as DELETES the repo from your system.",
		Action: func(_ ctx.Context, c *cli.Command) error {
			cmd := exec.Cmd(fmt.Sprintf("%s uninstall", cfg.Binary), exec.Stdio)

			// Uninstall the packages.
			assert.Will(cmd.Run())

			cmd = exec.Cmd(fmt.Sprintf("stow -t %s -D .build", home), exec.Stdio)
			cmd.Dir = cfg.Base

			// Uninstall the CLI.
			assert.Will(cmd.Run())

			// Delete the build.
			assert.Will(os.RemoveAll(cfg.Build))

			// Delete the repo locally.
			assert.Will(os.RemoveAll(cfg.Base))
			return nil
		},
	}
}
