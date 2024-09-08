package mappers

import (
	"users/internal/address"
	"users/pkg/adapters/storage/entities"
	"users/pkg/fp"
)

func AddressEntityToDomain(entity entities.Address) address.Address {
	return address.Address{ 
		UserID: entity.UserID,
		Street: entity.Street,
		City: entity.City,
		State: entity.State,
		ZipCode: entity.ZipCode,
		Country: entity.Country,
	}
}

func BatchAddressEntityToDomain(entities []entities.Address) []address.Address {
	return fp.Map(entities,AddressEntityToDomain)
}

func AddressDomainToEntity(domainaddress address.Address) entities.Address {
	return entities.Address{
		UserID: domainaddress.UserID,
		Street: domainaddress.Street,
		City: domainaddress.City,
		State: domainaddress.State,
		ZipCode: domainaddress.ZipCode,
		Country: domainaddress.Country,
	}
}