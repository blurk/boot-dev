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

### Ideas for Extending the Project

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

## HTTP CRUD Database Lifecycle

Here's how we could engineer GET request:

First, the front-end webpage loads.
The front-end sends an HTTP GET request to a /users endpoint on the back-end server.
The server receives the request.
The server uses a SELECT statement to retrieve the user's record from the users table in the database.
The server converts the row of SQL data into a JSON object and sends it back to the front-end.

## WHERE

You can use a WHERE clause to filter values by whether or not they're NULL.

### IS NULL
```SELECT name FROM users WHERE first_name IS NULL;```

### IS NOT NULL
```SELECT name FROM users WHERE first_name IS NOT NULL;```

## DELETING STRATEGY

### Strategy 1 - Backups
If you're using a cloud-service like GCP's Cloud SQL or AWS's RDS you should always turn on automated backups. They take an automatic snapshot of your entire database on some interval, and keep it around for some length of time.

For example, the Boot.dev database has a backup snapshot taken daily and we retain those backups for 30 days. If I ever accidentally run a query that deletes valuable data, I can restore it from the backup.

You should have a backup strategy for production databases.

### Strategy 2 - Soft Deletes
A "soft delete" is when you don't actually delete data from your database, but instead just "mark" the data as deleted. For example, you might set a deleted_at date on the row you want to delete. Then, in your queries you ignore anything that has a deleted_at date set. The idea is that this allows your application to behave as if it's deleting data, but you can always go back and restore any data that's been removed.

## Object-Relational Mapping (ORMs)
An Object-Relational Mapping or an ORM for short, is a tool that allows you to perform CRUD operations on a database using a traditional programming language.
These typically come in the form of a library or framework that you would use in your backend code.

The primary benefit an ORM provides is that it maps your database records to in-memory objects.