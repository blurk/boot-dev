package main

import "fmt"

func (e email) cost() int {
	costPerChar := 5

	if e.isSubscribed {
		costPerChar = 2
	}

	return len(e.body) * costPerChar
}

func (e email) format() string {
	endString := "Not Subscribed"

	if e.isSubscribed {
		endString = "Subscribed"
	}

	return fmt.Sprintf("'%s' | %s", e.body, endString)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
