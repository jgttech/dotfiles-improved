package rebuild

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "rebuild",
		Usage: "Uninstall the packages, uninstall the CLI, delete the repo, and rebuild and reinstall everything from the source code.",
		Action: func(ctx context.Context, c *cli.Command) error {
			return nil
		},
	}
}
