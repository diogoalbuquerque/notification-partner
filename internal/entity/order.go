package entity

type Order struct {
	Code           int64   `json:"code" example:"22"`
	NrChip         int64   `json:"nrChip" example:"3670549064"`
	SeqOnlineOrder string  `json:"nrSeqOnlineOrder" example:"100012912"`
	Token          string  `json:"token" example:"f7278f1c-c001-4790-bec9-6908b1a7da40"`
	VlCharge       float64 `json:"vlCharge" example:"60.00"`
	Cellphone      string  `json:"cellphone" example:"21998876655"`
}
