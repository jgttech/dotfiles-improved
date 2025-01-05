INSTALL_SCRIPT = .dotfiles/bin/install
CLI = $(shell bin/binary)

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

.PHONY: config
config:
	@podman exec -it dotfiles /bin/zsh -c "cat .dotfiles.build.json"

#######################################################
# Lifecycle commands
#

.PHONY: install
install: clean build init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true bash"

#######################################################
# Install methods
#
# Used for testing different installation processes for
# whichever language I want to use.

.PHONY: go
go: clean build init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=go bash"

.PHONY: odin
odin: clean build init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=odin bash"

.PHONY: zig
zig: clean build init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=zig bash"

.PHONY: python
python: clean build init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=python bash"

.PHONY: typescript
typescript: clean build init
	@podman exec -it dotfiles /bin/zsh -c \
		"cat $(INSTALL_SCRIPT) | DOTFILES_DEV=true DOTFILES_LANG=typescript bash"
