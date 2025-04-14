package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	result := make([][]byte, len(items))

	for i, item := range items {
		data, err := json.Marshal(item)

		if err != nil {
			return nil, err
		}

		result[i] = data
	}

	return result, nil
}
