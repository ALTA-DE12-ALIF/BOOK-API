package delivery

import (
	"main.go/features/user/domain"
)

type RegisterFormat struct {
	Nama     string `json:"nama" form:"nama"`
	HP       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

type LoginFormat struct {
	HP       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Nama: cnv.Nama, HP: cnv.HP, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{HP: cnv.HP, Password: cnv.Password}
	}

	return domain.Core{}
}
