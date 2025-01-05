package env

import (
	ctx "context"
	"fmt"
	"jgttech/dotfiles/src/context"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v3"
)

func Command(etx *context.ExecutionContext) *cli.Command {
	cfg := etx.Build.Config
	title := lipgloss.NewStyle().Bold(true).Faint(true).Render

	return &cli.Command{
		Name:  "env",
		Usage: "",
		Action: func(ctx ctx.Context, cmd *cli.Command) error {
			variable := cmd.Args().First()

			if variable == "" {
				var sb strings.Builder

				sb.WriteString("\n")
				sb.WriteString(title("[BUILD CONFIGURATION]"))
				sb.WriteString("\n")

				sb.WriteString("base=" + cfg.Base + "\n")
				sb.WriteString("build=" + cfg.Build + "\n")
				sb.WriteString("tools=" + cfg.Tools + "\n")
				sb.WriteString("binary=" + cfg.Binary + "\n")
				sb.WriteString("symlink=" + cfg.Symlink + "\n")
				sb.WriteString("language=" + cfg.Language + "\n")
				sb.WriteString("source=" + cfg.Source + "\n\n")

				sb.WriteString(title("[PACKAGES]"))
				sb.WriteString("\n")
				for _, pkg := range etx.Packages() {
					sb.WriteString(pkg + "\n")
				}
				sb.WriteString("\n")

				sb.WriteString(title("[TOOLS]"))
				sb.WriteString("\n")
				for _, tool := range etx.Tools() {
					sb.WriteString(tool + "\n")
				}

				fmt.Println(sb.String())
				return nil
			}

			value := reflect.ValueOf(*cfg)
			types := value.Type()

			for idx := 0; idx < value.NumField(); idx++ {
				prop := strings.ToLower(types.Field(idx).Name)

				if prop == variable {
					fmt.Println(value.Field(idx).Interface())
					return nil
				}
			}

			fmt.Println(fmt.Sprintf("Unable to find variable: '%s'", variable))
			return nil
		},
	}
}
