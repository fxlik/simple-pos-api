package web

import "time"

type QtyTypeCreateRequest struct {
	Name      string    `validate:"required,min=1,max=255"  json:"name" form:"name"`
	Disabled  bool      `json:"disabled" form:"disabled"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type QtyTypeUpdateRequest struct {
	Id        int32     `validate:"required" json:"id" form:"id"`
	Name      string    `validate:"required" json:"name" form:"name"`
	Disabled  bool      `json:"disabled" form:"disabled"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
