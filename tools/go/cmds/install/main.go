package install

import (
	"context"
	"fmt"
	// "fmt"
	// "jgttech/dotfiles/src/assert"
	// "jgttech/dotfiles/src/exec"
	// "os"

	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "install",
		Usage: "",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("INSTALLING")
			// command := fmt.Sprintf("stow -t %s", os.Getenv("HOME"))
			// cmd := exec.Cmd(command, exec.Stdio)

			return nil
		},
	}
}
