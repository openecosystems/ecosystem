# Setup

## Install Homebrew

`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)”`

## Install GoLand

## Install dependencies and configure terminal

```bash
brew install buf \
fastly/tap/fastly \
protobuf@3 \
hugo \
golangci-lint \
goreleaser/tap/goreleaser-pro \
goreleaser-pro \
ko \
go \
git \
automake \
gettext \
ghz \
diffutils \
httpie \
jq \
pnpm@8 \
telnet \
wget \
nvm
```

Run the following:
`mkdir ~/.nvm`
`echo 'export PATH="/opt/homebrew/opt/pnpm@8/bin:$PATH"' >> ~/.zshrc`
`nvm install 20.9.0`

Add the following to your `.zshrc`:

```bash
export PATH="/opt/homebrew/opt/pnpm@8/bin:$PATH"
export NVM_DIR="$HOME/.nvm"
[ -s "/opt/homebrew/opt/nvm/nvm.sh" ] && \. "/opt/homebrew/opt/nvm/nvm.sh"  # This loads nvm
[ -s "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm" ] && \. "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion
```

Run the following commands to configure Go tooling:
`echo 'export GOPATH="~/go"'`
`echo 'export PATH="$PATH:$GOROOT/bin:$GOPATH/bin"'`

Install Go dependencies:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
go install mvdan.cc/gofumpt@latest
go install golang.org/x/vuln/cmd/govulncheck@latest
```

## Install Rust

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y --default-toolchain 1.81.0
rustup install 1.81.0
rustup target add wasm32-wasi --toolchain 1.81.0
```

## IDE Setup

Run steps from [IDE](./IDE.md) file

## Install project dependencies

Install dependencies with pnpm `pnpm install`

## Generate SSH key

-   Follow the instructions here to add your SSH key https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent
-   Add the following line to the top of your .zshrc:
    `ssh-add --apple-use-keychain ~/.ssh/id_ed25519 /dev/null 2>&1`
-   Create a config file: `touch ~/.ssh/config`
-   Add the following to the config file:

```bash
UserKnownHostsFile ~/.ssh/known_hosts
Host *
  AddKeysToAgent yes
  UseKeychain yes
  IdentityFile ~/.ssh/id_ed25519
```

-   Add your public key from `~./ssh/id_ed25519.pub` to GitHub
-   Add your known host with `ssh -v git@github.com` and enter `yes` when prompted
-   Run the following in the project parent directory to change git remote to use SSH instead of HTTP: `set-url origin git@github.com:openecosystems/ecosystem.git`
