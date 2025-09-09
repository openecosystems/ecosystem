# Goland

## Configure Goland GO Settings

-   File -> Settings -> Go -> GOROOT: go version 1.23.11
-   File -> Settings -> Go -> GOPATH; Ensure GOPATH is correct
-   File -> Settings -> Go -> Go Modules: Enable Go Modules Support

## Install the following Plugins

-   https://plugins.jetbrains.com/plugin/12496-go-linter
-   https://plugins.jetbrains.com/plugin/7391-asciidoc
-   https://plugins.jetbrains.com/plugin/19147-buf-for-protocol-buffers
-   https://plugins.jetbrains.com/plugin/10275-hunspell
-   https://plugins.jetbrains.com/plugin/20146-mermaid
-   https://plugins.jetbrains.com/plugin/21060-nx-console

## Enable Golangcli-lint File Watcher

-   File -> Settings -> Tools -> File Watchers
-   -   sign
-   "golangcli-lint" option
-   Save

# Remove default import sorting type. Will be managed by golangcli-lint

-   File -> Settings -> Editor -> Code Style -> Go -> Imports -> Sorting Type: None

# Enable Gofumpt (a stricter gofmt)

GoLand doesn't use gopls so it should be configured to use gofumpt directly. Once gofumpt is installed, follow the steps below:

Open Settings (File > Settings)
Open the Tools section
Find the File Watchers sub-section
Click on the + on the right side to add a new file watcher
Choose Custom Template
When a window asks for settings, you can enter the following:

File Types: Select all .go files
Scope: Project Files
Program: Select your gofumpt executable
Arguments: -w $FilePath$
Output path to refresh: $FilePath$
Working directory: $ProjectFileDir$
Environment variables: GOROOT=$GOROOT$;GOPATH=$GOPATH$;PATH=$GoBinDirs$
To avoid unnecessary runs, you should disable all checkboxes in the Advanced section.

# Remove Actions on Save to prevent multiple formatting conflicts

-   File -> Settings -> Tools -> Actions on Save -> Deselect Reformat Code
