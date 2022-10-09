package repository

import (
	"main.go/features/pinjam/domain"

	"gorm.io/gorm"
)

type Pinjam struct {
	gorm.Model
	ID_buku         uint
	ID_peminjam     uint
	Tanggal_pinjam  string
	Tanggal_kembali string
}

func FromDomain(du domain.Core) Pinjam {
	return Pinjam{
		Model:           gorm.Model{ID: du.ID},
		ID_buku:         du.ID_buku,
		ID_peminjam:     du.ID_peminjam,
		Tanggal_pinjam:  du.Tanggal_pinjam,
		Tanggal_kembali: du.Tanggal_kembali,
	}
}

func ToDomain(u Pinjam) domain.Core {
	return domain.Core{

		ID:              u.ID,
		ID_buku:         u.ID_buku,
		ID_peminjam:     u.ID_peminjam,
		Tanggal_pinjam:  u.Tanggal_pinjam,
		Tanggal_kembali: u.Tanggal_kembali,
	}
}

func ToDomainArray(au []Pinjam) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, ID_buku: val.ID_buku, ID_peminjam: val.ID_peminjam, Tanggal_pinjam: val.Tanggal_pinjam, Tanggal_kembali: val.Tanggal_kembali})
	}

	return res
}
