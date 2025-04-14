package main

func getLast[T any](s []T) T {
	var myZero T

	numElements := len(s)

	if numElements > 0 {
		return s[numElements-1]
	}

	return myZero
}
