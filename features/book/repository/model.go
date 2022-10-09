package repository

import (
	"main.go/features/book/domain"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul     string
	Pengarang string
	Pemilik   uint
}

func FromDomain(du domain.Core) Book {
	return Book{
		Model:     gorm.Model{ID: du.ID},
		Judul:     du.Judul,
		Pengarang: du.Pengarang,
		Pemilik:   du.Pemilik,
	}
}

func ToDomain(u Book) domain.Core {
	return domain.Core{
		ID:        u.ID,
		Judul:     u.Judul,
		Pengarang: u.Pengarang,
		Pemilik:   u.Pemilik,
	}
}

func ToDomainArray(au []Book) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Judul: val.Judul, Pengarang: val.Pengarang, Pemilik: val.Pemilik})
	}
	return res
}
