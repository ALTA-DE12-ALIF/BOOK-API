package repository

import (
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

func (rq *repoQuery) Insert(newBook domain.Core) (domain.Core, error) {
	var cnv Book
	cnv = FromDomain(newBook)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newBook = ToDomain(cnv)
	return newBook, nil
}

func (rq *repoQuery) Update(updatedBook domain.Core) (domain.Core, error) {
	var cnv Book
	cnv = FromDomain(updatedBook)
	if err := rq.db.Save(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	updatedBook = ToDomain(cnv)
	return updatedBook, nil
}
func (rq *repoQuery) Get(ID uint) (domain.Core, error) {
	var resQry Book
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}
func (rq *repoQuery) GetAll() ([]domain.Core, error) {
	var resQry []Book
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}
