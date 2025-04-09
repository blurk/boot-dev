package main

import "unicode"

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	hasUppercase := false
	hasDigit := false

	for _, c := range password {
		if unicode.IsUpper(c) {
			hasUppercase = true
		}

		if unicode.IsDigit(c) {
			hasDigit = true
		}

		// Early exit
		if hasDigit && hasUppercase {
			return true
		}
	}

	return hasDigit && hasUppercase
}
