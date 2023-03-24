package web

import "time"

type SupplierCreateRequest struct {
	Name      string    `validate:"required,min=1,max=255" json:"name"`
	Email     string    `validate:"required,min=1,max=255" json:"email"`
	Telp      string    `validate:"required,min=4,max=20" json:"telp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SupplierUpdateRequest struct {
	Id        int32     `json:"id"`
	Name      string    `validate:"required,min=1,max=255" json:"name"`
	Email     string    `validate:"required,min=1,max=255" json:"email"`
	Telp      string    `validate:"required,min=1,max=255" json:"telp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
