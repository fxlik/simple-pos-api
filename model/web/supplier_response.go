package web

import (
	"fxlik/simple-post-api/model/domain"
	"time"
)

type SupplierResponse struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Telp      string    `json:"telp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToSupplierResponse(supplier domain.Supplier) SupplierResponse {
	return SupplierResponse{
		Id:        supplier.Id,
		Name:      supplier.Name,
		Email:     supplier.Email,
		Telp:      supplier.Telp,
		CreatedAt: supplier.CreatedAt,
		UpdatedAt: supplier.UpdatedAt,
	}
}

func ToSupplierResponses(suppliers []domain.Supplier) []SupplierResponse {
	var supplierResponses []SupplierResponse
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, ToSupplierResponse(supplier))
	}
	return supplierResponses
}
