package client

import (
	"fmt"

	"github.com/SPVJ/fs-common-lib/core/client"
	"github.com/thanavatC/bff-document-service-go/config"
	"github.com/thanavatC/bff-document-service-go/model"
)

type CompanyServiceClient interface {
	GetCompanies() ([]*model.Company, error)
}

type CompanyServiceClientImpl struct {
	httpClient client.IHttpClient
}

func NewCompanyServiceClientImpl(httpClient client.IHttpClient) CompanyServiceClient {
	return &CompanyServiceClientImpl{
		httpClient: httpClient,
	}
}

func (c *CompanyServiceClientImpl) GetCompanies() ([]*model.Company, error) {
	var response []*model.Company

	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.GetCompanies
	url := fmt.Sprintf("%v%v/%v", base, path)
	headers := map[string]string{}

	if err := c.httpClient.Get(&response, url, headers); err != nil {
		return nil, err
	}

	return response, nil
}
