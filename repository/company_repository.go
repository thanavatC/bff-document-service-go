package repository

import (
	"github.com/SPVJ/document-service-go/model"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	GetCompanies() ([]model.Company, error)
}

type CompanyRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanyRepositoryImpl(db *gorm.DB) CompanyRepository {
	return &CompanyRepositoryImpl{
		db: db,
	}
}

func (repo *CompanyRepositoryImpl) GetCompanies() ([]model.Company, error) {
	var companies []model.Company
	if err := repo.db.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}
