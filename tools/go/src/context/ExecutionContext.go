package context

import (
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/config"
	"os"
	"path"
	"sync"
)

type ExecutionContext struct {
	Build    *config.BuildJson
	packages []string
	tools    []string
}

func NewExecutionContext() *ExecutionContext {
	var wg sync.WaitGroup
	build := config.NewBuildJson()
	packages := []string{}
	tools := []string{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		toolsDir := path.Join(build.Config.Base, "tools")
		toolsFiles := assert.Must(os.ReadDir(toolsDir))

		for _, file := range toolsFiles {
			tools = append(tools, file.Name())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		packagesDir := path.Join(build.Config.Base, "packages")
		packagesFiles := assert.Must(os.ReadDir(packagesDir))

		for _, file := range packagesFiles {
			packages = append(packages, file.Name())
		}
	}()

	wg.Wait()

	return &ExecutionContext{build, packages, tools}
}

func (etx *ExecutionContext) Packages() []string {
	return etx.packages
}

func (etx *ExecutionContext) Tools() []string {
	return etx.tools
}
