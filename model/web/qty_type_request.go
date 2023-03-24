package web

import "time"

type QtyTypeCreateRequest struct {
	Name      string    `validate:"required, min=1, max=255" json:"name"`
	Disabled  bool      `json:"disabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QtyTypeUpdateRequest struct {
	Id        int32     `validate:"required" json:"id"`
	Name      string    `validate:"required" json:"name"`
	Disabled  bool      `json:"disabled"`
	UpdatedAt time.Time `json:"updated_at"`
}
