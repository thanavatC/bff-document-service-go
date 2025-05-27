package client

import (
	"fmt"

	"github.com/SPVJ/fs-common-lib/core/client"
	"github.com/thanavatC/bff-document-service-go/config"
	"github.com/thanavatC/bff-document-service-go/model"
)

type DocumentRequestServiceClient interface {
	CreateDocumentRequest(req model.CreateDocumentRequestRequest) (*model.DocumentRequest, error)
	DeleteDocumentRequest(id string) error
	ListDocumentRequests(page, pageSize int) (*model.DocumentRequestListResponse, error)
	ValidateDocumentRequest(id string, req model.ValidateDocumentRequestRequest) (*model.DocumentRequest, error)
}

type DocumentRequestServiceClientImpl struct {
	httpClient client.IHttpClient
}

func NewDocumentRequestServiceClientImpl(httpClient client.IHttpClient) DocumentRequestServiceClient {
	return &DocumentRequestServiceClientImpl{
		httpClient: httpClient,
	}
}

func (c *DocumentRequestServiceClientImpl) CreateDocumentRequest(req model.CreateDocumentRequestRequest) (*model.DocumentRequest, error) {
	var response model.DocumentRequest

	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.CreateDocumentRequest
	url := fmt.Sprintf("%v%v", base, path)
	headers := map[string]string{}

	if err := c.httpClient.Post(&response, req, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *DocumentRequestServiceClientImpl) DeleteDocumentRequest(id string) error {
	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.DeleteDocumentRequest
	url := fmt.Sprintf("%v%v/%v", base, path, id)
	headers := map[string]string{}

	return c.httpClient.Delete(nil, url, headers)
}

func (c *DocumentRequestServiceClientImpl) ListDocumentRequests(page, pageSize int) (*model.DocumentRequestListResponse, error) {
	var response model.DocumentRequestListResponse

	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.ListDocumentRequests
	url := fmt.Sprintf("%v%v?page=%d&page_size=%d", base, path, page, pageSize)
	headers := map[string]string{}

	if err := c.httpClient.Get(&response, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *DocumentRequestServiceClientImpl) ValidateDocumentRequest(id string, req model.ValidateDocumentRequestRequest) (*model.DocumentRequest, error) {
	var response model.DocumentRequest

	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.ValidateDocumentRequest
	url := fmt.Sprintf("%v%v/%v/validate", base, path, id)
	headers := map[string]string{}

	if err := c.httpClient.Post(&response, req, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}
