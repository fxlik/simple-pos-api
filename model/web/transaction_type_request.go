package web

type TransactionTypeCreateRequest struct {
	Name     string `validate:"required" json:"name"`
	Disabled bool   `json:"disabled"`
}

type TransactionTypeUpdateRequest struct {
	Id       int32  `validate:"required" json:"id"`
	Name     string `validate:"required" json:"name"`
	Disabled bool   `json:"disabled"`
}
