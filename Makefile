.PHONY: image
image:
	@podman build -t dotfiles --format docker .

.PHONY: ssh
ssh:
	@podman run -it \
		-v $(shell pwd):/root/.dotfiles \
		--name dotfiles \
		dotfiles /bin/zsh

.PHONY: install
install:
	@podman run -t \
		-v $(shell pwd):/root/.dotfiles \
		--name dotfiles \
		dotfiles /bin/zsh -c "bash .dotfiles/bin/dev.install"

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
rebuild: nuke image
