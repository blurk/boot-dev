package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// Create a nil slice of items []Issue.
	var issues []Issue

	jsonDecoder := json.NewDecoder(res.Body)

	if err := jsonDecoder.Decode(&issues); err != nil {
		return nil, fmt.Errorf("error decoding response body")
	}

	return issues, nil

}
