package web

import "time"

type SupplierCreateRequest struct {
	Name      string    `validate:"required,min=1,max=255" json:"name" form:"name"`
	Email     string    `validate:"required,min=1,max=255" json:"email" form:"email"`
	Telp      string    `validate:"required,min=4,max=20" json:"telp" form:"telp"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type SupplierUpdateRequest struct {
	Id        int32     `json:"id" form:"id"`
	Name      string    `validate:"required,min=1,max=255" json:"name" form:"name"`
	Email     string    `validate:"required,min=1,max=255" json:"email" form:"email"`
	Telp      string    `validate:"required,min=1,max=255" json:"telp" form:"telp"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
