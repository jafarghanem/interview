package address

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	CreateAddress(ctx context.Context, address *Address)( *Address ,error)
	GetAddress(ctx context.Context,userid uuid.UUID, page, pageSize int) ([]Address, uint, error)
}

type Address struct {
	ID      uint
	UserID	uuid.UUID
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}
