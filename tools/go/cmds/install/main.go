package install

import (
	ctx "context"
	"fmt"
	"jgttech/dotfiles/cmds/install/alacritty"
	"jgttech/dotfiles/cmds/install/ghostty"
	"jgttech/dotfiles/cmds/install/tpm"
	"jgttech/dotfiles/src/context"
	"jgttech/dotfiles/src/exec"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli/v3"
)

func Command(etx *context.ExecutionContext) *cli.Command {
	cfg := etx.Build.Config
	home := os.Getenv("HOME")

	return &cli.Command{
		Name:  "install",
		Usage: "Install all the packages.",
		Action: func(_ ctx.Context, c *cli.Command) error {
			args := fmt.Sprintf("stow -t %s %s", home, strings.Join(etx.Packages(), " "))
			cmd := exec.Cmd(args, exec.Stdio)
			cmd.Dir = path.Join(cfg.Base, "packages")

			if err := cmd.Run(); err != nil {
				return err
			}

			tpm.Install()
			alacritty.Install(etx)
			ghostty.Install(etx)

			return nil
		},
	}
}
