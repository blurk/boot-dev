package main

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))

	for _, email := range emails {
		emailsToSend <- email
	}

	return emailsToSend
}
