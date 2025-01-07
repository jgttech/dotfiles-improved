package save

import (
	ctx "context"
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/context"
	"jgttech/dotfiles/src/exec"
	"strconv"
	"time"

	"github.com/urfave/cli/v3"
)

func Command(etx *context.ExecutionContext) *cli.Command {
	cfg := etx.Build.Config

	return &cli.Command{
		Name:  "save",
		Usage: "Commit and changes to the repo.",
		Action: func(_ ctx.Context, c *cli.Command) error {
			msg := c.Args().First()

			if msg == "" {
				msg = "Updated (" + strconv.Itoa(int(time.Now().Unix())) + ")"
			}

			commands := []string{
				"git add .",
				"git commit -m '" + msg + "'",
				"git push",
			}

			for _, command := range commands {
				cmd := exec.Cmd(command, exec.Stdio)
				cmd.Dir = cfg.Base

				assert.Will(cmd.Run())
			}

			return nil
		},
	}
}
