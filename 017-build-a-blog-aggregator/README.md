## Guide

1. Install postgress
2. Install go
3. Run `go build .` to build or just run `go run .` in the root

## Set up config file in the root

```bash
touch ~/.gatorconfig.json && echo "{\"db_url\":\"postgres://username:password@localhost:5432/db_name?sslmode=disable\"}" > ~/.gatorconfig.json
```

## Command list and usage

- To login: `login <name>`
- To register: `register <name>`
- To reset: `reset`
- To see user list: `users`
- To crawl data posts: `agg <duration between request>`
- To add new feed: `addfeed <name> <url>`
- To list all feeds: `feeds`
- To follow: `follow <url>`
- To show following feeds of current user: `following`
- To unfollow a feed: `unfollow <url>`
- To see all the post of the following feeds: `browse`