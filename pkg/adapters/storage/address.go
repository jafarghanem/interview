package storage

import (
	"context"
	"errors"
	"users/internal/address"
	"users/pkg/adapters/storage/entities"
	"users/pkg/adapters/storage/mappers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) address.Repo {
	return &addressRepo{
		db: db,
	}
}


func (r *addressRepo) CreateAddress(ctx context.Context, a *address.Address) (*address.Address,error) {
	addressEntity := mappers.AddressDomainToEntity(*a)
	if err := r.db.WithContext(ctx).Create(&addressEntity).Error; err != nil {
		return nil ,err
	}
	addressDomain := mappers.AddressEntityToDomain(addressEntity)
	ad := &addressDomain
	return ad,nil
}

func (r *addressRepo) GetAddress(ctx context.Context, userid uuid.UUID, page, pageSize int) ([]address.Address, uint, error) {
	var a []entities.Address
	var int64Total int64

	// Filter addresses by UserID and preload related Users
	query := r.db.Model(&entities.Address{}).
		Where("user_id = ?", userid). // Filter by userid
		Preload("Users")

	// Count total records for pagination
	query.Count(&int64Total)

	// Pagination logic
	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// Execute the query
	if err := query.Find(&a).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	// Convert total records to uint
	total := uint(int64Total)

	// Map entities to domain objects
	addresses := mappers.BatchAddressEntityToDomain(a)

	return addresses, total, nil
}
