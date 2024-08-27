package model

// NOT_CONFIRMED, PENDING, PROCESSED, CANCELLED, FAILED, DONE
type PaymentStatus string

type Payment struct {
	Id             string          `json:"id,omitempty"`
	LinkId         string          `json:"link_id,omitempty"`
	Status         PaymentStatus   `json:"status,omitempty"`
	PaymentChannel *PaymentChannel `json:"payment_channel,omitempty"`
}

type PaymentChannel struct {
	AccountNumber string `json:"account_number,omitempty"`
	AccountType   string `json:"account_type,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
	AccountHolder string `json:"account_holder,omitempty"`
	QrCodeData    string `json:"qr_code_data,omitempty"`
}
