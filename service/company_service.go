package service

import (
	"github.com/SPVJ/document-service-go/model"
	"github.com/SPVJ/document-service-go/repository"
)

type CompanyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

func (s *CompanyService) GetCompanies() ([]model.Company, error) {
	return s.repo.GetCompanies()
}
