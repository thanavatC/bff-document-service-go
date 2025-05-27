package service

import (
	"github.com/thanavatC/bff-document-service-go/client"
	"github.com/thanavatC/bff-document-service-go/model"
)

type CompanyService interface {
	GetCompanies() ([]*model.Company, error)
}

type CompanyServiceImpl struct {
	companyServiceClient client.CompanyServiceClient
}

func NewCompanyServiceImpl(client client.CompanyServiceClient) CompanyService {
	return &CompanyServiceImpl{
		companyServiceClient: client,
	}
}

func (s *CompanyServiceImpl) GetCompanies() ([]*model.Company, error) {
	return s.companyServiceClient.GetCompanies()
}
