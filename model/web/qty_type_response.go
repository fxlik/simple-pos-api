package web

import (
	"fxlik/simple-post-api/model/domain"
	"time"
)

type QtyTypeResponse struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Disabled  bool      `json:"disabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToQtyTypeResponse(qtyType domain.QtyType) QtyTypeResponse {
	return QtyTypeResponse{
		Id:        qtyType.Id,
		Name:      qtyType.Name,
		Disabled:  qtyType.Disabled,
		CreatedAt: qtyType.CreatedAt,
		UpdatedAt: qtyType.UpdatedAt,
	}
}

func ToQtyTypeResponses(qtyTypes []domain.QtyType) []QtyTypeResponse {
	var qtyTypeResponses []QtyTypeResponse
	for _, qtyType := range qtyTypes {
		qtyTypeResponses = append(qtyTypeResponses, ToQtyTypeResponse(qtyType))
	}
	return qtyTypeResponses
}
