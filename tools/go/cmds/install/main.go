package install

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

	return &cli.Command{
		Name:  "install",
		Usage: "Link all the stow packages and run post-stow installation operations.",
		Action: func(ctx context.Context, c *cli.Command) error {
			stow := []string{}
			dir := path.Join(cfg.Base, "packages")
			packages := assert.Must(os.ReadDir(dir))

			for _, pkg := range packages {
				stow = append(stow, pkg.Name())
			}

			cmd := exec.Cmd(
				fmt.Sprintf(
					"stow -t %s %s",
					os.Getenv("HOME"),
					strings.Join(stow, " "),
				),
				exec.Stdio,
			)

			cmd.Dir = dir

			return cmd.Run()
		},
	}
}
