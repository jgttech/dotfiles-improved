package reinstall

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "reinstall",
		Usage: "Uninstall the packages, rebuild the CLI, and re-install the packages.",
		Action: func(ctx context.Context, c *cli.Command) error {
			return nil
		},
	}
}
