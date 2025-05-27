package client

import (
	"fmt"

	"github.com/SPVJ/fs-common-lib/core/client"
	"github.com/thanavatC/bff-document-service-go/config"
	"github.com/thanavatC/bff-document-service-go/model"
)

type DocumentServiceClient interface {
	GetDocumentStatus(id string) (string, error)
	ReTranslateDocument(id string) (*model.Document, error)
	UpdateDocument(id string, req model.UpdateDocumentRequest) (*model.Document, error)
	DeleteDocument(id string) error
	ValidateDocument(id string, req model.ValidateDocumentRequest) (*model.Document, error)
}

type DocumentServiceClientImpl struct {
	httpClient client.IHttpClient
}

func NewDocumentServiceClientImpl(httpClient client.IHttpClient) DocumentServiceClient {
	return &DocumentServiceClientImpl{
		httpClient: httpClient,
	}
}

func (c *DocumentServiceClientImpl) GetDocumentStatus(id string) (string, error) {
	var response struct {
		Status string `json:"status"`
	}

	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.GetDocumentStatus
	url := fmt.Sprintf("%v%v/%v", base, path, id)
	headers := map[string]string{}

	if err := c.httpClient.Get(&response, url, headers); err != nil {
		return "", err
	}

	return response.Status, nil
}

func (c *DocumentServiceClientImpl) ReTranslateDocument(id string) (*model.Document, error) {
	var response model.Document

	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.ReTranslateDocument
	url := fmt.Sprintf("%v%v/%v", base, path, id)
	headers := map[string]string{}

	if err := c.httpClient.Post(&response, nil, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *DocumentServiceClientImpl) UpdateDocument(id string, req model.UpdateDocumentRequest) (*model.Document, error) {
	var response model.Document

	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.UpdateDocument
	url := fmt.Sprintf("%v%v/%v", base, path, id)
	headers := map[string]string{}

	if err := c.httpClient.Patch(&response, req, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *DocumentServiceClientImpl) DeleteDocument(id string) error {
	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.DeleteDocument
	url := fmt.Sprintf("%v%v/%v", base, path, id)
	headers := map[string]string{}

	return c.httpClient.Delete(nil, url, headers)
}

func (c *DocumentServiceClientImpl) ValidateDocument(id string, req model.ValidateDocumentRequest) (*model.Document, error) {
	var response model.Document

	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.ValidateDocument
	url := fmt.Sprintf("%v%v/%v", base, path, id)
	headers := map[string]string{}

	if err := c.httpClient.Post(&response, req, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}
