package models

// MailData represents the input data schema for sending emails.
type MailData struct {
	Sender     string   `json:"sender"`
	Subject    string   `json:"subject"`
	HTML       string   `json:"html"`
	Text       string   `json:"text"`
	Recipients []string `json:"recipients"`
}
