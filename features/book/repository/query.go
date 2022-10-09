package repository

import (
	"errors"

	"main.go/features/book/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) InsertBook(newBook domain.Core) (domain.Core, error) {
	var cnv Book
	cnv = FromDomain(newBook)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	newBook = ToDomain(cnv)
	return newBook, nil
}

func (rq *repoQuery) UpdateBook(updatedData domain.Core) (domain.Core, error) {
	var cnv Book
	cnv = FromDomain(updatedData)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	updatedData = ToDomain(cnv)
	return updatedData, nil
}

func (rq *repoQuery) GetAllBook() ([]domain.Core, error) {
	var resQry []Book
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) DeleteBook(ID domain.Core) error {
	var res Book = FromDomain(ID)
	if err := rq.db.Where("id = ?", res.ID).Delete(&res).Error; err != nil {
		return errors.New("gagal delete")
	}
	// ID = ToDomain(res)
	return errors.New("berhasil delete")
}
