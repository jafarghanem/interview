package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"users/api/http/handlers"
	"users/api/http/middlewares"
	"users/config"
	"users/service"
)

func Run(cfg config.Config, app *service.AppContainer) {
	fiberApp := fiber.New()

	api := fiberApp.Group("/api/v1", middlewares.SetUserContext())

	// register global routes
	registerGlobalRoutes(api, app)
	secret := []byte(cfg.Server.TokenSecret)
	registerAddressRoutes(api, app, secret)
	log.Fatal(fiberApp.Listen(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)))
}


func registerGlobalRoutes(router fiber.Router, app *service.AppContainer) {
	router.Post("/register", handlers.RegisterUser(app.AuthService()))
	router.Post("/registerconc", handlers.RegisterUserConcurent(app.AuthService()))
	router.Post("/login", handlers.LoginUser(app.AuthService()))
	router.Get("/refresh", handlers.RefreshToken(app.AuthService()))
}


func registerAddressRoutes(router fiber.Router, app *service.AppContainer, secret []byte) {
	router.Post("/address",middlewares.Auth(secret),handlers.CreateAddress(app.AddressService()))
	router.Post("/addressconc", handlers.CreateAddressConc(app.AddressService()))
	router.Get("/my-address",middlewares.Auth(secret),handlers.GetAddresss(app.AddressService()))

}


