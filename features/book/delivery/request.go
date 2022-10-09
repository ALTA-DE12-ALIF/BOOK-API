package delivery

import (
	"main.go/features/book/domain"
)

type RegisterFormat struct {
	Judul     string `json:"judul" form:"judul"`
	Pengarang string `json:"pengarang" form:"pengarang"`
	Pemilik   string `json:"pemilik" form:"pemilik"`
}

type LoginFormat struct {
	Pengarang string `json:"pengarang" form:"pengarang"`
	Pemilik   string `json:"pemilik" form:"pemilik"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Judul: cnv.Judul, Pengarang: cnv.Pengarang, Pemilik: cnv.Pemilik}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Pengarang: cnv.Pengarang, Pemilik: cnv.Pemilik}
	}

	return domain.Core{}
}
