package main

import "fmt"

func fetchTasks(baseURL, availability string) []Issue {
	qSort := "estimate"
	qLimit := ""

	switch availability {
	case "Low":
		qLimit += "1"
	case "Medium":
		qLimit += "3"
	case "High":
		qLimit += "5"
	}

	fullURL := baseURL + fmt.Sprintf("?sort=%v&limit=%v", qSort, qLimit)

	return getIssues(fullURL)
}
