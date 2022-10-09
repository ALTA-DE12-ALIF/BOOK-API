package services

import (
	"errors"
	"strings"

	"main.go/features/pinjam/domain"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type pinjamService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &pinjamService{
		qry: repo,
	}
}

func (us *pinjamService) AddPinjam(newPinjam domain.Core) (domain.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newPinjam.Tanggal_pinjam), bcrypt.DefaultCost)

	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encript password")
	}

	newPinjam.Tanggal_pinjam = string(generate)
	res, err := us.qry.Insert(newPinjam)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}

		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

func (us *pinjamService) UpdateProfile(updatePinjam domain.Core) (domain.Core, error) {
	if updatePinjam.Tanggal_pinjam != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(updatePinjam.Tanggal_pinjam), bcrypt.DefaultCost)
		if err != nil {
			log.Error(err.Error())
			return domain.Core{}, errors.New("cannot encript password")
		}
		updatePinjam.Tanggal_pinjam = string(generate)
	}

	res, err := us.qry.Update(updatePinjam)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
	}

	return res, nil
}
func (us *pinjamService) Profile(ID uint) (domain.Core, error) {
	res, err := us.qry.Get(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}

	return res, nil
}
func (us *pinjamService) ShowAllPinjam() ([]domain.Core, error) {
	res, err := us.qry.GetAll()
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}

	return res, nil
}
