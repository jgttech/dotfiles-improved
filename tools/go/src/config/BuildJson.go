package config

import (
	"encoding/json"
	"fmt"
	"io"
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/env"
	"jgttech/dotfiles/src/os"
	_os "os"
	"path"
)

type ConfigJson struct {
	Base     string `json:"base"`
	Language string `json:"language"`
	Tools    string `json:"tools"`
	Build    string `json:"build"`
	Binary   string `json:"binary"`
	Symlink  string `json:"symlink"`
	Source   string `json:"source"`
}

type BuildJson struct {
	Config *ConfigJson `json:"config"`
}

func NewBuildJson() *BuildJson {
	buildFile := path.Join(_os.Getenv("HOME"), env.DOTFILES_BUILD_CONFIG)

	if !os.Exists(buildFile) {
		panic(fmt.Sprintf("Build config not found: '%s'", buildFile))
	}

	file := assert.Must(_os.Open(buildFile))
	defer file.Close()

	bytes := assert.Must(io.ReadAll(file))
	build := &BuildJson{&ConfigJson{}}

	json.Unmarshal([]byte(bytes), build)

	return build
}
