package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	encoder := json.NewDecoder(res.Body)

	var users []User

	if err := encoder.Decode(&users); err != nil {
		return nil, err
	}

	return users, nil

}
