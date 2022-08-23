package entity

type Notification struct {
	ClientRechargeToken  string `json:"clientRechargeToken" example:"f7278f1c-c001-4790-bec9-6908b1a7da40"`
	PaymentRequestNumber string `json:"paymentRequestNumber" example:"123"`
	ClientIdentity       string `json:"clientIdentity" example:"21998876655"`
}
