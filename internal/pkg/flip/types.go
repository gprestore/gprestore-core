package flip

import "time"

const (
	FlipBillTypeSingle   = "SINGLE"
	FlipBillTypeMultiple = "MULTIPLE"

	FlipCodeBNI          = "bni"
	FlipCodeBRI          = "bri"
	FlipCodeBCA          = "bca"
	FlipCodeMandiri      = "mandiri"
	FlipCodeCIMB         = "cimb"
	FlipCodeBTPN         = "tabungan_pensiunan_nasional"
	FlipCodeDBS          = "dbs"
	FlipCodePermata      = "permata"
	FlipCodeMuamalat     = "muamalat"
	FlipCodeDanamon      = "danamon"
	FlipCodeBSI          = "bsm"
	FlipCodeOvo          = "ovo"
	FlipCodeQris         = "qris"
	FlipCodeShopeepayApp = "shopeepay_app"
	FlipCodeLinkAja      = "linkaja"
	FlipCodeLinkAjaApp   = "linkaja_app"
	FlipCodeDana         = "dana"

	FlipAccountTypeBankAccount            = "bank_account"
	FlipAccountTypeVirtualAccount         = "virtual_account"
	FlipAccountTypeWalletAccount          = "wallet_account"
	FlipAccountTypeOnlineToOfflineAccount = "online_to_offline_account"
	FlipAccountTypeCreditCardAccount      = "credit_card_account"
)

type FlipBillRequest struct {
	Title                 string     `json:"title,omitempty"`
	Type                  string     `json:"type,omitempty"`
	Amount                int        `json:"amount,omitempty"`
	RedirectUrl           string     `json:"redirect_url,omitempty"`
	ExpiredDate           *time.Time `json:"expired_date,omitempty"`
	CreatedFrom           string     `json:"created_from,omitempty"`
	Status                string     `json:"status,omitempty"`
	IsAddressRequired     int        `json:"is_address_required,omitempty"`
	IsPhoneNumberRequired int        `json:"is_phone_number_required,omitempty"`
	Step                  string     `json:"step,omitempty"`
	SenderName            string     `json:"sender_name,omitempty"`
	SenderEmail           string     `json:"sender_email,omitempty"`
	SenderPhoneNumber     string     `json:"sender_phone_number,omitempty"`
	SenderAddress         string     `json:"sender_address,omitempty"`
	SenderBank            string     `json:"sender_bank,omitempty"`
	SenderBankType        string     `json:"sender_bank_type,omitempty"`
}

type FlipBill struct {
	LinkId                int                `json:"link_id,omitempty"`
	LinkUrl               string             `json:"link_url,omitempty"`
	Title                 string             `json:"title,omitempty"`
	Type                  string             `json:"type,omitempty"`
	Amount                int                `json:"amount,omitempty"`
	RedirectUrl           string             `json:"redirect_url,omitempty"`
	ExpiredDate           *time.Time         `json:"expired_date,omitempty"`
	CreatedFrom           string             `json:"created_from,omitempty"`
	Status                string             `json:"status,omitempty"`
	IsAddressRequired     int                `json:"is_address_required,omitempty"`
	IsPhoneNumberRequired int                `json:"is_phone_number_required,omitempty"`
	Step                  int                `json:"step,omitempty"`
	Customer              *FlipCustomer      `json:"customer,omitempty"`
	BillPayment           *FlipBillPayment   `json:"bill_payment,omitempty"`
	PaymentMethod         *FlipPaymentMethod `json:"payment_method,omitempty"`
	PaymentUrl            string             `json:"payment_url,omitempty"`
}

type FlipCustomer struct {
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

type FlipBillPayment struct {
	Id                  string                   `json:"id,omitempty"`
	Amount              int                      `json:"amount,omitempty"`
	UniqueCode          int                      `json:"unique_code,omitempty"`
	Status              string                   `json:"status,omitempty"`
	SenderBank          string                   `json:"sender_bank,omitempty"`
	SenderBankType      string                   `json:"sender_bank_type,omitempty"`
	ReceiverBankAccount *FlipReceiverBankAccount `json:"receiver_bank_account,omitempty"`
	UserAddress         string                   `json:"user_address,omitempty"`
	UserPhone           string                   `json:"user_phone,omitempty"`
	CreatedAt           int64                    `json:"created_at,omitempty"`
}

type FlipReceiverBankAccount struct {
	AccountNumber string `json:"account_number,omitempty"`
	AccountType   string `json:"account_type,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
	AccountHolder string `json:"account_holder,omitempty"`
	QrCodeData    string `json:"qr_code_data,omitempty"`
}

type FlipPaymentMethod struct {
	SenderBank     string `json:"sender_bank,omitempty"`
	SenderBankType string `json:"sender_bank_type,omitempty"`
}

type FlipPayment struct {
	LinkId      int               `json:"link_id,omitempty"`
	TotalData   int               `json:"total_data,omitempty"`
	DataPerPage int               `json:"data_per_page,omitempty"`
	TotalPage   int               `json:"total_page,omitempty"`
	Page        int               `json:"page,omitempty"`
	Data        []FlipPaymentData `json:"data,omitempty"`
}

type FlipPaymentData struct {
	Id                   string `json:"id,omitempty"`
	LinkId               string `json:"link_id,omitempty"`
	BillLink             string `json:"bill_link,omitempty"`
	BillTitle            string `json:"bill_title,omitempty"`
	SenderName           string `json:"sender_name,omitempty"`
	SenderBank           string `json:"sender_bank,omitempty"`
	SenderBankType       string `json:"sender_bank_type,omitempty"`
	VirtualAccountNumber string `json:"virtual_account_number,omitempty"`
	Amount               int    `json:"amount,omitempty"`
	Status               string `json:"status,omitempty"`
	SettlementStatus     string `json:"settlement_status,omitempty"`
	ReferenceId          string `json:"reference_id,omitempty"`
	PaymentUrl           string `json:"payment_url,omitempty"`
	CreatedAt            string `json:"created_at,omitempty"`
	CompletedAt          string `json:"completed_at,omitempty"`
	SettlementDate       string `json:"settlement_date,omitempty"`
}
