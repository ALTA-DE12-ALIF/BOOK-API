package delivery

import "main.go/features/book/domain"

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
	ID        uint   `json:"id"`
	Judul     string `json:"Judul"`
	Pengarang string `json:"pengarang"`
	Pemilik   uint   `json:"pemilik"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, Judul: cnv.Judul, Pengarang: cnv.Pengarang, Pemilik: cnv.Pemilik}
	case "all":
		var arr []RegisterResponse
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, RegisterResponse{ID: val.ID, Judul: val.Judul, Pengarang: val.Pengarang, Pemilik: val.Pemilik})
		}
		res = arr
	case "del":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, Judul: cnv.Judul, Pengarang: cnv.Pengarang, Pemilik: cnv.Pemilik}
	case "update":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, Judul: cnv.Judul, Pengarang: cnv.Pengarang, Pemilik: cnv.Pemilik}
	}
	return res
}
