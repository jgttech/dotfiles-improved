package config

import (
	"fmt"
	"jgttech/dotfiles/src/assert"
	"os"
	"path"
	"strings"
)

func (build *BuildJson) StowUninstall() string {
	home := os.Getenv("HOME")
	stow := []string{}
	dir := path.Join(build.Config.Base, "packages")
	packages := assert.Must(os.ReadDir(dir))

	for _, pkg := range packages {
		stow = append(stow, pkg.Name())
	}

	return fmt.Sprintf("stow -t %s -D %s", home, strings.Join(stow, " "))
}
