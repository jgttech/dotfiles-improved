package packages

import (
	ctx "context"
	"fmt"
	"jgttech/dotfiles/src/context"
	"strings"

	"github.com/urfave/cli/v3"
)

func Command(etx *context.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:  "packages",
		Usage: "Display what packages currently exist.",
		Action: func(_ ctx.Context, c *cli.Command) error {
			fmt.Println(strings.Join(etx.Packages(), "\n"))
			return nil
		},
	}
}
