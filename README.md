- To run `create_folder.sh`: run with `. [path to the script]` to make sure cd work

- Cool packages: curl and jq (files in /tmp/ will be removed after 10days)
 ```bash
curl https://api.boot.dev/v1/courses_rest_api/learn-http/issues | jq '.[].title, .[].estimate' > /tmp/issue_info.txt
 ```
 ```bash
curl -X POST https://api.boot.dev/v1/courses_rest_api/learn-http/users -H "Content-Type: application/json" -d '{
  "role": "QA Job Safety",
  "experience": 2,
  "remote": true,
  "user": {
    "name": "Dan",
    "location": "NOR",
    "age": 29
  }
}' > /tmp/user.json
 ```

### Ideas for extending the Project: 015 Pokedex

- [x] Update the CLI to support the "up" arrow to cycle through previous commands
- [ ] Simulate battles between pokemon
- [ ] Add more unit tests
- [ ] Refactor your code to organize it better and make it more testable
- [ ] Keep pokemon in a "party" and allow them to level up
- [ ] Allow for pokemon that are caught to evolve after a set amount of time
- [ ] Persist a user's Pokedex to disk so they can save progress between sessions
- [ ] Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- [ ] Random encounters with wild pokemon
- [ ] Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon

### Ideas for extending the Project: 017 Blog aggregator

- [ ] Add sorting and filtering options to the browse command
- [ ] Add pagination to the browse command
- [ ] Add concurrency to the agg command so that it can fetch more frequently
- [ ] Add a search command that allows for fuzzy searching of posts
- [ ] Add bookmarking or liking posts
- [ ] Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)
- [ ] Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
- [ ] Write a service manager that keeps the agg command running in the background and restarts it if it crashes