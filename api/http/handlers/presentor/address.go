package presenter

import (
	"users/internal/address"

	"github.com/google/uuid"
)

type CreateAddressReq struct {
	UserID  uuid.UUID `json:"user_id"`
	Street  string    `json:"street" validate:"required"`
	City    string    `json:"city" validate:"required"`
	State   string    `json:"state" validate:"required"`
	ZipCode string    `json:"zip_code" validate:"required"`
	Country string    `json:"country" validate:"required"`
}
type CreateAddressConcReq struct {
	UserID  uuid.UUID `json:"user_id" validate:"required"`
	Street  string    `json:"street" validate:"required"`
	City    string    `json:"city" validate:"required"`
	State   string    `json:"state" validate:"required"`
	ZipCode string    `json:"zip_code" validate:"required"`
	Country string    `json:"country" validate:"required"`
}
type AddressResp struct {
	ID      uint      `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Street  string    `json:"street"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	ZipCode string    `json:"zip_code"`
	Country string    `json:"country"`
}

func CreateAddressConcRequest(req *CreateAddressConcReq) *address.Address {
	return &address.Address{
		UserID:  req.UserID,
		Street:  req.Street,
		City:    req.City,
		State:   req.State,
		ZipCode: req.ZipCode,
		Country: req.Country,
	}
}
func CreateAddressRequest(req *CreateAddressReq) *address.Address {
	return &address.Address{
		UserID:  req.UserID,
		Street:  req.Street,
		City:    req.City,
		State:   req.State,
		ZipCode: req.ZipCode,
		Country: req.Country,
	}
}

func AddressToCreateAddressResponse(r *address.Address) *AddressResp {
	return &AddressResp{
		ID:      r.ID,
		UserID:  r.UserID,
		Street:  r.Street,
		City:    r.City,
		State:   r.State,
		ZipCode: r.ZipCode,
		Country: r.Country,
	}
}

func AddressToFullAddressResponse(r *address.Address) *AddressResp {
	return AddressToCreateAddressResponse(r)
}
