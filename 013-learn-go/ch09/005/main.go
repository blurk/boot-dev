package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, messagedUser := range messagedUsers {
		if _, isValidUser := validUsers[messagedUser]; isValidUser {
			validUsers[messagedUser]++
		}
	}
}
