# Personal Dotfiles Ecosystem

Because I am insane and love to program for fun. I decided that each just having some dotfiles is not enough for me. I wanted my dotfiles to also be a strange place where I can install varients of those tools written in different languages, as I see fit, whenever I feel like it. This is just to have some fun with my own tools and also gain some experience in other languages building CLI-based applications.

# Installation

This invokes the install script and builds the built-in CLI.

```bash
wget -qO- "https://raw.githubusercontent.com/jgttech/dotfiles/refs/heads/main/bin/install" | bash
```

# Prerequisites

These are requied to do development with this repo.

# Development

The development of the CLI uses `podman` to create a container and install the CLI within the container. Because linked volumes within containers can affect the underlying operating system, no volumes are used. All development is basically done by creating a container that contains the installed CLI and then running commands against it and inspecting the file system within the container to see the results of the command without breaking the host OS.

> To setup the development environment run:

```bash
bun i
```

> Create a container with an installation.

```bash
# Alternatively, you can also use "make install".
# Afterall, this just invokes "make" anyway.
bun +install
```

## Development Commands

Commands can be invoked to run against the installed CLI within the container, once it is setup. Creating the container copies over the files and installs the CLI to allow things to be tested without breaking the host OS.

| Command | Description |
|:-|:-|
| `+nuke` | Destroys all containers and images. |
| `+clean` | Destroys all containers. |
| `+ssh` | Connect to the container. |
| `+config` | Cat the build configuration. |
| `bun +install` | Builds/Rebuilds the container and runs the default installation. |
| `bun +go` | Builds/Rebuilds the container and runs the Go language installation. |
| `bun +odin` | Builds/Rebuilds the container and runs the Odin language installation. |
| `bun +zig` | Builds/Rebuilds the container and runs the Zig language installation. |
| `bun +python` | Builds/Rebuilds the container and runs the Python language installation. |
| `bun +typescript` | Builds/Rebuilds the container and runs the TypeScript language installation. |
| `bun @go ...` | Run script against the workspace for Go lang. |
| `bun @odin ...` | Run script against the workspace for Odin lang. |
| `bun @zig ...` | Run script against the workspace for Zig lang. |
| `bun @python ...` | Run script against the workspace for Python lang. |
| `bun @typescript ...` | Run script against the workspace for TypeScript lang. |
