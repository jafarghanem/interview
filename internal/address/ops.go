package address

import (
	"context"

	"github.com/google/uuid"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo}
}

func (o *Ops) CreateAddress(ctx context.Context, Address *Address) (*Address,error) {
	return o.repo.CreateAddress(ctx, Address)
}

func (o *Ops) GetAddress(ctx context.Context,userid uuid.UUID, page, pageSize int) ([]Address, uint, error) {
	return o.repo.GetAddress(ctx,userid, page, pageSize)
}
