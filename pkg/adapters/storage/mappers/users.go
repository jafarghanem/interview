package mappers

import (
	"users/internal/user"
	"users/pkg/adapters/storage/entities"
	"users/pkg/fp"
)

func UserEntityToDomain(entity entities.User) user.User {
	domainAddress := BatchAddressEntityToDomain(entity.Address)
	return user.User{
		ID : entity.ID,
		FirstName:   entity.FirstName,
		LastName: entity.LastName,
		Email:  entity.Email,
		Password: entity.Password,
		PhoneNumber: entity.PhoneNumber,
		Address: domainAddress,
	}
}
func userEntityToDomain(entity entities.User) user.User {
	return user.User{
		ID : entity.ID,
		FirstName: entity.FirstName,
		LastName: entity.LastName,
		Email: entity.Email,
		PhoneNumber: entity.PhoneNumber,
	}
}

func BatchUserEntityToDomain(entities []entities.User) []user.User {
	return fp.Map(entities, userEntityToDomain)
}

func UserDomainToEntity(domainUser user.User) entities.User {
	return entities.User{
		ID : domainUser.ID,
		FirstName: domainUser.FirstName,
		LastName:  domainUser.LastName,
		Email:     domainUser.Email,
		Password:  domainUser.Password,
		PhoneNumber: domainUser.PhoneNumber,
	}
}