
# Build
go build -o oeco
go build -o $GOPATH/bin/oeco
GOOS=windows go build -o oeco.exe


# Publish Dry Run
goreleaser --snapshot --skip-publish --rm-dist

# Publish
goreleaser release --skip-publish


# Manually publish release until automated

- Commit all changes in the repository, make sure nothing is "dirty"

- Create a tag
``sh
git tag 0.0.0
git push origin : 0.0.0
``

- Export AWS Profile
``sh
export AWS_PROFILE=sf-us-west-2-operations
``

- Export Fake GITHUB token
``sh
export GITHUB_TOKEN=0
``

- Update your credentials
``sh
sf login

select operations admin role
``

- Run goreleaser
``sh
goreleaser release --rm-dist
``

- Update the homebrew-tap repository sf.rb file

Get the checksum from the dist/checksums.txt file and update the macos checksum in the sf.rb file
Also increment the version number to reflect the tag version

- Commit the homebrew-tap repo

- To delete a tag
``sh
git tag -d 0.0.0
git push origin : refs/tags/0.0.0
``

Installation Instructions

# Assumptions:
- Homebrew is installed
- SSH key is configure to talk to git@bitbucket.org:jeannotcompany

# Steps
brew tap jeannotcompany/homebrew-tap git@bitbucket.org:jeannotcompany/homebrew-tap.git
brew install sf


# If something goes wrong
brew uninstall sf
brew untap jeannotcompany/homebrew-tap
brew tap jeannotcompany/homebrew-tap git@bitbucket.org:jeannotcompany/homebrew-tap.git
brew install sf

# To upgrade versino
brew update
brew upgrade sf







# Install

We have provided several pre-compiled binaries for easy installation,
or you can compile from source.


Below are the steps to install:

## Install the pre-compiled binary

**homebrew tap** (only on macOS for now):

```sh
brew tap jeannotcompany/homebrew-tap git@bitbucket.org:jeannotcompany/homebrew-tap.git
brew install sf
```

**homebrew** (may not be the latest version):

```sh
brew install sf-sdk-cli
```

**scoop**:

```powershell
Set-ExecutionPolicy RemoteSigned -scope CurrentUser
iwr -useb get.scoop.sh | iex
---

scoop bucket add sf https://github.com/jeannotcompany/scoop-bucket.git
scoop bucket add sf git@bitbucket.org:jeannotcompany/scoop-sf.git
scoop install sf
```

**manually**:

Download the pre-compiled binaries from the [releases page][releases] and
copy to the desired location.


[releases]: https://github.com/jeannotcompany/sf-sdk-cli/releases

## Compiling from source

If you want to build from source, follow these steps:

**Clone:**

```sh
git clone https://bitbucket.org/jeannotcompany/sf-sdk-cli
cd sf-sdk-cli
```

**Get the dependencies:**

```sh
go get ./...
```

**Build:**

```sh
go build -o sf .
```

**Verify it works:**

```sh
./sf --version
```
