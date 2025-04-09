package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	userPhoneMap := map[string]user{}

	for i := range names {
		userPhoneMap[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
	}

	return userPhoneMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
