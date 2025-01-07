package edit

import (
	ctx "context"
	"jgttech/dotfiles/src/context"
	"jgttech/dotfiles/src/exec"
	"path"

	"github.com/urfave/cli/v3"
)

func Command(etx *context.ExecutionContext) *cli.Command {
	cfg := etx.Build.Config

	return &cli.Command{
		Name:  "edit",
		Usage: "Opens the source code for the currently installed tools within Neovim.",
		Action: func(_ ctx.Context, c *cli.Command) error {
			tools := path.Join(cfg.Base, cfg.Tools)
			return exec.Cmd("nvim "+tools, exec.Stdio).Run()
		},
	}
}
