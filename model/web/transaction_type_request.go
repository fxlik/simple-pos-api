package web

type TransactionTypeCreateRequest struct {
	Name     string `validate:"required" json:"name" form:"name"`
	Disabled bool   `json:"disabled" form:"disabled"`
}

type TransactionTypeUpdateRequest struct {
	Id       int32  `validate:"required" json:"id" form:"id"`
	Name     string `validate:"required" json:"name" form:"name"`
	Disabled bool   `json:"disabled" form:"disabled"`
}
