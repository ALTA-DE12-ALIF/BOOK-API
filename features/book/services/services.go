package services

import (
	"errors"
	"strings"

	"main.go/features/book/domain"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type bookService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &bookService{
		qry: repo,
	}
}

func (bs *bookService) AddBook(newBook domain.Core) (domain.Core, error) {

	res, err := bs.qry.InsertBook(newBook)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

func (bs *bookService) UpdateBook(updatedData domain.Core) (domain.Core, error) {
	res, err := bs.qry.UpdateBook(updatedData)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
	}
	return res, nil
}

func (bs *bookService) ShowAllBook() ([]domain.Core, error) {
	res, err := bs.qry.GetAllBook()
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("No data")
		}
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}

	return res, nil
}

func (bs *bookService) Delete(ID domain.Core) error {
	err := bs.qry.DeleteBook(ID)
	if strings.Contains(err.Error(), "") {
		return errors.New("cant find the data")
	}
	return errors.New("berhasil delete")
}

func (bs *bookService) ExtractToken(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		return uint(claim["id"].(float64))
	}

	return 0
}
