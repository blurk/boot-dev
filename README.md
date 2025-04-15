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