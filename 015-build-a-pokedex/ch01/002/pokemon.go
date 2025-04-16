package main

type ResponseLocation struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"` // Can be null
	Results  []Location `json:"results"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
