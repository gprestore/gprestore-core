package model

type Mail struct {
	From    string   `json:"from,omitempty"`
	To      []string `json:"to,omitempty"`
	Cc      *MailCc  `json:"cc,omitempty"`
	Subject string   `json:"subject,omitempty"`
	Body    string   `json:"body,omitempty"`
}

type MailCc struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
}
