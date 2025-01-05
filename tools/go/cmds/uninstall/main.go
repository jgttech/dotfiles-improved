package uninstall

import (
	"context"
	"fmt"
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/config"
	"jgttech/dotfiles/src/exec"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli/v3"
)

func Command(build *config.BuildJson) *cli.Command {
	cfg := build.Config
	home := os.Getenv("HOME")

	return &cli.Command{
		Name:  "uninstall",
		Usage: "Unlink all the stow packages. But the CLI will still exist.",
		Action: func(ctx context.Context, c *cli.Command) error {
			// Detect and unlink all the packages.
			stow := []string{}
			packagesDir := path.Join(cfg.Base, "packages")
			packages := assert.Must(os.ReadDir(packagesDir))

			for _, pkg := range packages {
				stow = append(stow, pkg.Name())
			}

			cmd := exec.Cmd(
				fmt.Sprintf(
					"stow -t %s -D %s",
					home,
					strings.Join(stow, " "),
				),
				exec.Stdio,
			)

			cmd.Dir = packagesDir
			assert.Will(cmd.Run())

			return nil
		},
	}
}
