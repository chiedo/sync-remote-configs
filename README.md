# Sync Remote Configs (beta)

A tool to keep your remote users in sync with your local configurations for things like vim and git.

1. Install or updated by running this

    `wget -O /usr/local/bin/sync-remote-configs https://github.com/chiedo/sync-remote-configs/raw/master/sync-remote-configs && chmod +x /usr/local/bin/sync-remote-configs`
    
2. Create a `~/.sync-remote-configs/destinations` file in the repo
    - Each line must be in `username@domain` format
3. Create an `~/.sync-remote-configs/exclusions` file in the repo.
    - Each line is a relative file path such as `.git/*`
4. Create a `~/.sync-remote-configs/sources` file in the repo.
    - Each line is an absolute path to a directory you want synced with each destination.
5. Run `sync-remote-configs`


# Development

- Test by making changes to `main.go` and running `go run main.go`
- Build by running `go build`
