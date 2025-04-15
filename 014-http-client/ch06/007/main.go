package main

import (
	"net/http"
)

func getUserCode(url string) int {
	res, err := http.Get(url)
	if err != nil {
		return res.StatusCode
	}
	defer res.Body.Close()

	return res.StatusCode
}
