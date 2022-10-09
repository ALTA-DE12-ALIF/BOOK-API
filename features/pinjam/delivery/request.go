package delivery

import (
	"main.go/features/pinjam/domain"
)

type PinjamFormat struct {
	ID_buku         uint   `json:"id_buku" form:"id_buku"`
	ID_peminjam     uint   `json:"id_peminjam" form:"id_peminjam"`
	Tanggal_pinjam  string `json:"tanggal_pinjam" form:"tanggal_pinjam"`
	Tanggal_kembali string `json:"tanggal_kembali" form:"tanggal_kembali"`
}

type KembaliFormat struct {
	ID_peminjam     uint   `json:"id_peminjam" form:"id_peminjam"`
	Tanggal_pinjam  string `json:"tanggal_pinjam" form:"tanggal_pinjam"`
	Tanggal_kembali string `json:"tanggal_kembali" form:"tanggal_kembali"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case PinjamFormat:
		cnv := i.(PinjamFormat)
		return domain.Core{ID_buku: cnv.ID_buku, ID_peminjam: cnv.ID_peminjam, Tanggal_pinjam: cnv.Tanggal_pinjam, Tanggal_kembali: cnv.Tanggal_kembali}
	case KembaliFormat:
		cnv := i.(KembaliFormat)
		return domain.Core{ID_peminjam: cnv.ID_peminjam, Tanggal_pinjam: cnv.Tanggal_pinjam, Tanggal_kembali: cnv.Tanggal_kembali}
	}

	return domain.Core{}
}
