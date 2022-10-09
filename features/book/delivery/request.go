package delivery

import (
	"main.go/features/book/domain"
)

type RegisterFormat struct {
	Judul     string `json:"judul" form:"judul"`
	Pengarang string `json:"pengarang" form:"pengarang"`
	Pemilik   uint   `json:"pemilik" fomr:"pemilik"`
}

type UpdateFormat struct {
	ID        uint   `json:"id" form:"id"`
	Judul     string `json:"judul" form:"judul"`
	Pengarang string `json:"pengarang" form:"pengarang"`
	Pemilik   uint   `json:"pemilik" fomr:"pemilik"`
}

type MyBookFormat struct {
	Pemilik uint `json:"pemilik" fomr:"pemilik"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Judul: cnv.Judul, Pengarang: cnv.Pengarang, Pemilik: cnv.Pemilik}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, Judul: cnv.Judul, Pengarang: cnv.Pengarang, Pemilik: cnv.Pemilik}
	case MyBookFormat:
		cnv := i.(MyBookFormat)
		return domain.Core{Pemilik: cnv.Pemilik}
	}
	return domain.Core{}
}
