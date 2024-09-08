package service

import (
	"context"
	"users/internal/address"
	"users/internal/user"

	"github.com/google/uuid"
)

type AddressService struct {
	userOps    *user.Ops
	addressOps *address.Ops
}

func NewAddressService(userOps *user.Ops, addressOps *address.Ops) *AddressService {
	return &AddressService{
		userOps:    userOps,
		addressOps: addressOps,
	}
}

func (s *AddressService) CreateAddress(ctx context.Context, a *address.Address) (*address.Address,error) {
	return s.addressOps.CreateAddress(ctx, a)
}

func (s *AddressService) GetAddress(ctx context.Context,userid uuid.UUID, page, pageSize int) ([]address.Address, uint, error) {
	return s.addressOps.GetAddress(ctx,userid, page, pageSize)
}