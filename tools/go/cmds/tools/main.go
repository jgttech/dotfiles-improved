package tools

import (
	ctx "context"
	"fmt"
	"jgttech/dotfiles/src/context"
	"strings"

	"github.com/urfave/cli/v3"
)

func Command(etx *context.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:  "tools",
		Usage: "Display what tools (languages) currently exist.",
		Action: func(_ ctx.Context, c *cli.Command) error {
			fmt.Println(strings.Join(etx.Tools(), "\n"))
			return nil
		},
	}
}
