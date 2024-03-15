package types

type Email struct {
	To      []string
	Subject string
	Body    string
}

// NewEmail creates a new Email object with the specified recipients, subject, and body.
func NewEmail(to []string, subject, body string) Email {
	return Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}
