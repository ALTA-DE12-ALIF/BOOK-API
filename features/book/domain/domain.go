package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	Judul     string
	Pengarang string
	Pemilik   uint
}

type Repository interface {
	InsertBook(newBook Core) (Core, error)
	UpdateBook(newBook Core) (Core, error)
	GetAllBook() ([]Core, error)
	DeleteBook(ID Core) error
}

type Service interface {
	AddBook(newBook Core) (Core, error)
	UpdateBook(updatedData Core) (Core, error)
	ShowAllBook() ([]Core, error)
	Delete(ID Core) error
	ExtractToken(c echo.Context) uint
}

type Handler interface {
	AddBook() echo.HandlerFunc
	ShowAllBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
}
