package web

type TransactionStatusCreateRequest struct {
	Name     string `validate:"required" json:"name"`
	Disabled bool   `json:"disabled"`
}

type TransactionStatusUpdateRequest struct {
	Id       int32  `validate:"required" json:"id"`
	Name     string `validate:"required" json:"name"`
	Disabled bool   `json:"disabled"`
}
