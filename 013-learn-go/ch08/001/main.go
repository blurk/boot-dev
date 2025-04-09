package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	primaryLen := len(primary)
	secondaryLen := len(secondary)
	tertiaryLen := len(tertiary)

	return [3]string{primary, secondary, tertiary}, [3]int{primaryLen, primaryLen + secondaryLen, primaryLen + secondaryLen + tertiaryLen}
}
