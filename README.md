# Sync Remote Configs (beta)

A tool to keep your remote users in sync with your local configurations for things like vim and git.

1. Clone this repo into your gopath
2. Create a `.destinations` file in the repo
    - Each line must be in `username@domain` format
3. Create an `.exclusions` file in the repo.
    - Each line is a relative file path such as `.git/*`
4. Create a `.sources` file in the repo.
    - Each line is an absolute path to a directory you want synced with each destination.
5. Run `go run main.go`