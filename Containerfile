FROM ubuntu:latest

WORKDIR /root

RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y \
  git \
  curl \
  wget \
  vim \
  neovim \
  zsh \
  jq \
  golang-go \
  stow \
  unzip \
  gnupg2 \
  lsb-release \
  sudo \
  build-essential \
  && rm -rf /var/lib/apt/lists/*;

RUN curl -fsSL https://bun.sh/install | bash;

# Set the shell to ZSH
SHELL ["/bin/zsh", "-c"]

# Add Bun to the PATH (this is critical for ZSH)
RUN echo 'export PATH="/root/.bun/bin:$PATH"' >> /root/.zshrc;

# Add Bun to the PATH
ENV PATH="/root/.bun/bin:$PATH"

RUN sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
