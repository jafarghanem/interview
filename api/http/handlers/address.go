package handlers

import (
	"users/api/http/handlers/presentor"
	"users/pkg/jwt"
	"users/service"

	"github.com/gofiber/fiber/v2"
)

func CreateAddress(addressService *service.AddressService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.CreateAddressReq

		
		if err := c.BodyParser(&req); err != nil {
			return presenter.BadRequest(c, err)
		}

		
		userClaims, ok := c.Locals(UserClaimKey).(*jwt.UserClaims)
		if !ok {
			return SendError(c, errWrongClaimType, fiber.StatusBadRequest)
		}

		req.UserID = userClaims.UserID

		err := BodyValidator(req)
		if err != nil {
			return presenter.BadRequest(c, err)
		}
		
		a := presenter.CreateAddressRequest(&req)

		
		createdAddress, err := addressService.CreateAddress(c.UserContext(), a)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		
		res := presenter.AddressToCreateAddressResponse(createdAddress)
		return presenter.Created(c, "Address created successfully", res)
	}
}
func CreateAddressConc(addressService *service.AddressService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req presenter.CreateAddressConcReq

		
		if err := c.BodyParser(&req); err != nil {
			return presenter.BadRequest(c, err)
		}


		err := BodyValidator(req)
		if err != nil {
			return presenter.BadRequest(c, err)
		}

		
		a := presenter.CreateAddressConcRequest(&req)

	
		createdAddress, err := addressService.CreateAddress(c.UserContext(), a)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

	
		res := presenter.AddressToCreateAddressResponse(createdAddress)
		return presenter.Created(c, "Address created successfully", res)
	}
}
func GetAddresss(addressService *service.AddressService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims, ok := c.Locals(UserClaimKey).(*jwt.UserClaims)
		if !ok {
			return presenter.BadRequest(c, errWrongClaimType)
		}
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("page_size", 10)

		Addresss, total, err := addressService.GetAddress(c.UserContext(),userClaims.UserID, page, pageSize)
		if err != nil {
			return presenter.InternalServerError(c, err)
		}

		res := make([]presenter.AddressResp, len(Addresss))
		for i, Address := range Addresss {
			res[i] = *presenter.AddressToFullAddressResponse(&Address) // Dereference the pointer here
		}

		pagination := presenter.NewPagination(res, uint(page), uint(pageSize), uint(total))
		return presenter.OK(c, "Addresss retrieved successfully", pagination)
	}
}