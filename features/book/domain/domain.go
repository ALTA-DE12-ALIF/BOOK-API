package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	Judul     string
	Pengarang string
	Pemilik   string
}

type Repository interface {
	Insert(newBook Core) (Core, error)
	Update(updateBook Core) (Core, error)
	Get(ID uint) (Core, error)
	GetAll() ([]Core, error)
}

type Service interface {
	AddBook(newBook Core) (Core, error)
	UpdateBook(updateBook Core) (Core, error)
	Profile(ID uint) (Core, error)
	ShowAllBook() ([]Core, error)
}

type Handler interface {
	AddBook() echo.HandlerFunc
	ShowAllBook() echo.HandlerFunc
}
