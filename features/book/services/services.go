package services

import (
	"errors"
	"strings"

	"main.go/features/book/domain"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type bookService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &bookService{
		qry: repo,
	}
}

func (us *bookService) AddBook(newBook domain.Core) (domain.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newBook.Judul), bcrypt.DefaultCost)

	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encript password")
	}

	newBook.Judul = string(generate)
	res, err := us.qry.Insert(newBook)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}

		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

func (us *bookService) UpdateBook(UpdateBook domain.Core) (domain.Core, error) {
	if UpdateBook.Judul != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(UpdateBook.Judul), bcrypt.DefaultCost)
		if err != nil {
			log.Error(err.Error())
			return domain.Core{}, errors.New("cannot encript judul")
		}
		UpdateBook.Judul = string(generate)
	}

	res, err := us.qry.Update(UpdateBook)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
	}

	return res, nil
}
func (us *bookService) Profile(ID uint) (domain.Core, error) {
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
func (us *bookService) ShowAllBook() ([]domain.Core, error) {
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
