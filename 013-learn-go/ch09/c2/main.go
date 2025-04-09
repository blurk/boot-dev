package main

import "slices"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	userFriends := friendships[username]
	suggestedFriends := []string{}

	for _, friend := range userFriends {
		friendsOfFriend := friendships[friend]

		for _, potentialFriend := range friendsOfFriend {
			if potentialFriend == username || slices.Contains(userFriends, potentialFriend) || slices.Contains(suggestedFriends, potentialFriend) {
				continue
			}

			suggestedFriends = append(suggestedFriends, potentialFriend)
		}
	}

	if len(suggestedFriends) != 0 {
		return suggestedFriends
	}

	return nil
}
