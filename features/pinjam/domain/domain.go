package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID              uint
	ID_buku         uint
	ID_peminjam     uint
	Tanggal_pinjam  string
	Tanggal_kembali string
}

type Repository interface {
	Insert(newPinjam Core) (Core, error)
	Update(updatePinjam Core) (Core, error)
	Get(ID uint) (Core, error)
	GetAll() ([]Core, error)
}

type Service interface {
	AddPinjam(newUser Core) (Core, error)
	UpdateProfile(updatePinjam Core) (Core, error)
	Profile(ID uint) (Core, error)
	ShowAllPinjam() ([]Core, error)
}

type Handler interface {
	AddPinjam() echo.HandlerFunc
	ShowAllPinjam() echo.HandlerFunc
}
