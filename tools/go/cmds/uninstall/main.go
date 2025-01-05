package uninstall

import (
	ctx "context"
	"fmt"
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
		Name:  "uninstall",
		Usage: "Uninstall the packages, but NOT the CLI.",
		Action: func(_ ctx.Context, c *cli.Command) error {
			args := fmt.Sprintf("stow -t %s -D %s", home, strings.Join(etx.Packages(), " "))
			cmd := exec.Cmd(args, exec.Stdio)
			cmd.Dir = path.Join(cfg.Base, "packages")

			return cmd.Run()
		},
	}
}
