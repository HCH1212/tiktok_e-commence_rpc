package model

type Email struct {
	From        string `json:"from"`
	To          string `json:"to"`
	ContentType string `json:"content_type"`
	Subject     string `json:"subject"`
	Content     string `json:"content"`
}
