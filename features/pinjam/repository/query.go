package repository

import (
	"main.go/features/pinjam/domain"

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

func (rq *repoQuery) Insert(newPinjam domain.Core) (domain.Core, error) {
	var cnv Pinjam
	cnv = FromDomain(newPinjam)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newPinjam = ToDomain(cnv)
	return newPinjam, nil
}

func (rq *repoQuery) Update(updatePinjam domain.Core) (domain.Core, error) {
	var cnv Pinjam
	cnv = FromDomain(updatePinjam)
	if err := rq.db.Save(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	updatePinjam = ToDomain(cnv)
	return updatePinjam, nil
}
func (rq *repoQuery) Get(ID uint) (domain.Core, error) {
	var resQry Pinjam
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}
func (rq *repoQuery) GetAll() ([]domain.Core, error) {
	var resQry []Pinjam
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}
