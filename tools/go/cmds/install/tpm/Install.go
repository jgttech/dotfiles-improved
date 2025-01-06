package tpm

import (
	"jgttech/dotfiles/src/assert"
	"jgttech/dotfiles/src/exec"
	"jgttech/dotfiles/src/os"
	_os "os"
	"path"
)

func Install() {
	base := ".tmux/plugins/tpm"
	dir := path.Join(_os.Getenv("HOME"), base)

	if os.Exists(dir) {
		return
	}

	clone := "git clone https://github.com/tmux-plugins/tpm "
	clone += dir

	assert.Will(exec.Cmd(clone, exec.Stdio).Run())
}
