package services

import (
	"errors"
	"strings"

	"main.go/features/user/domain"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

func (us *userService) AddUser(newUser domain.Core) (domain.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encript password")
	}

	newUser.Password = string(generate)
	res, err := us.qry.Insert(newUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}

		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

func (us *userService) UpdateProfile(updatedData domain.Core) (domain.Core, error) {
	if updatedData.Password != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error(err.Error())
			return domain.Core{}, errors.New("cannot encript password")
		}
		updatedData.Password = string(generate)
	}

	res, err := us.qry.Update(updatedData)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
	}

	return res, nil
}
func (us *userService) Profile(ID uint) (domain.Core, error) {
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
func (us *userService) ShowAllUser() ([]domain.Core, error) {
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
