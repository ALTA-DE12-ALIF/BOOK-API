package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Nama     string
	HP       string
	Password string
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	Update(updateData Core) (Core, error)
	Get(ID uint) (Core, error)
	GetAll() ([]Core, error)
}

type Service interface {
	AddUser(newUser Core) (Core, error)
	UpdateProfile(updateData Core) (Core, error)
	Profile(ID uint) (Core, error)
	ShowAllUser() ([]Core, error)
}

type Handler interface {
	AddUser() echo.HandlerFunc
	ShowAllUser() echo.HandlerFunc
}
