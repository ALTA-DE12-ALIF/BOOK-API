package delivery

import "main.go/features/pinjam/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type RegisterResponse struct {
	ID              uint   `json:"id"`
	ID_buku         uint   `json:"id_buku"`
	ID_peminjam     uint   `json:"id_peminjam"`
	Tanggal_pinjam  string `json:"tanggal_pinjam"`
	Tanggal_kembali string `json:"tanggal_kembali"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, ID_buku: cnv.ID_buku, ID_peminjam: cnv.ID_peminjam, Tanggal_pinjam: cnv.Tanggal_pinjam, Tanggal_kembali: cnv.Tanggal_kembali}
	case "all":
		var arr []RegisterResponse
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, RegisterResponse{ID: val.ID, ID_buku: val.ID_buku, ID_peminjam: val.ID_peminjam, Tanggal_pinjam: val.Tanggal_pinjam, Tanggal_kembali: val.Tanggal_kembali})
		}
		res = arr
	}

	return res
}
