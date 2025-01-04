# @podman run -dit \
# 	--name dotfiles \
# 	-v $(shell pwd):/root/.dotfiles \
# 	dotfiles
INSTALL_SCRIPT = .dotfiles/bin/install
UNINSTALL_SCRIPT = .dotfiles/bin/uninstall
REINSTALL_SCRIPT = .dotfiles/bin/reinstall

.PHONY: build
build:
	@podman build -t dotfiles --format docker .

.PHONY: init
init:
	@podman run -dit --name dotfiles dotfiles

.PHONY: ssh
ssh:
	@podman exec -it dotfiles /bin/zsh

.PHONY: nuke
nuke:
	-podman kill `podman ps -qa`
	-podman rm -f `podman ps -qa`
	-podman rmi -f `podman images -qa`

.PHONY: clean
clean:
	-podman kill `podman ps -qa`
	-podman rm -f `podman ps -qa`

.PHONY: rebuild
rebuild: nuke postinstall

.PHONY: reinstall
reinstall:
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(REINSTALL_SCRIPT) | DOTFILES_LANG=zig DOTFILES_DEV=true bash"

.PHONY: uninstall
uninstall:
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(UNINSTALL_SCRIPT) | DOTFILES_DEV=true bash"

# Default install method.
.PHONY: install
install: clean init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true bash"

# Go install method.
.PHONY: go
go: clean build init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=go bash"

# Odin install method.
.PHONY: odin
odin: clean init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=odin bash"

# Zig install method.
.PHONY: zig
zig: clean init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=zig bash"

# Python install method.
.PHONY: python
python: clean init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=python bash"

# TypeScript install method.
.PHONY: typescript
typescript:
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=typescript bash"
