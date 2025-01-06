package reinstall

import (
	ctx "context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "reinstall",
		Usage: "Uninstall the packages, CLI, and source code. Download the code again and run the installation script.",
		Action: func(_ ctx.Context, cmd *cli.Command) error {
			return nil
		},
	}
}
