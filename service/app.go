package service

import (
	"log"
	"users/config"
	"users/internal/address"
	"users/internal/user"
	"users/pkg/adapters/storage"

	"gorm.io/gorm"
)

type AppContainer struct {
	cfg            config.Config
	dbConn         *gorm.DB
	authService    *AuthService
	addressService *AddressService
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}

	app.mustInitDB()

	app.setAuthService()
	app.setAddressService()

	return app, nil
}

func (a *AppContainer) RawDBConnection() *gorm.DB {
	return a.dbConn
}

func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	db, err := storage.NewPostgresGormConnection(a.cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.dbConn = db

	err = storage.Migrate(a.dbConn)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
	err = storage.AddExtension(a.dbConn)
	if err != nil {
		log.Fatal("Create extension failed: ", err)
	}
}

func (a *AppContainer) AuthService() *AuthService {
	return a.authService
}

func (a *AppContainer) setAuthService() {
	if a.authService != nil {
		return
	}

	a.authService = NewAuthService(user.NewOps(storage.NewUserRepo(a.dbConn)), []byte(a.cfg.Server.TokenSecret),
		a.cfg.Server.TokenExpMinutes,
		a.cfg.Server.RefreshTokenExpMinutes)
}

func (a *AppContainer) AddressService() *AddressService {
	return a.addressService
}

func (a *AppContainer) setAddressService() {
	if a.addressService != nil {
		return
	}
	a.addressService = NewAddressService(user.NewOps(storage.NewUserRepo(a.dbConn)), address.NewOps(storage.NewAddressRepo(a.dbConn)))
}
