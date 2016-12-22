# Sync Remote Configs (beta)

A tool to keep your remote users in sync with your local configurations for things like vim and git.

1. Install or updated by running this

    `wget -O /usr/local/bin/sync-remote-configs https://github.com/chiedo/sync-remote-configs/raw/master/sync-remote-configs?date=$(date +%s) && chmod +x /usr/local/bin/sync-remote-configs`
    
2. Create a `~/.sync-remote-configs/destinations` file in the repo
    - Each line must be in `username@domain` format
    - You must have SSH access to each server via a public key.
    - These will be the servers that are getting updated.
3. Create a `~/.sync-remote-configs/exclusions` file in the repo.
    - Each line is a relative file path such as `.git/*`
    - This will be a list of file paths to ignore.
4. Create a `~/.sync-remote-configs/sources` file in the repo.
    - Each line is an absolute path to a directory you want synced with each destination. For now it's only been tested with directories in user home directories. Should work with other absolute file paths in theory.
    - This will be each directory that you intend to sync with your remote servers.
5. Run `sync-remote-configs`


# Development

- Test by making changes to `main.go` and running `go run main.go`
- Build by running `go build`

# Gotchas

- You have to have already successfully connected to a server before using this tool and if you get a known_host verification failure, things will break.
- I could turn off the strict host checking but it seems like the wrong move from a security perspective...
