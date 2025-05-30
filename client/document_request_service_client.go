package client

import (
	"fmt"
	"strings"

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
	var response []model.DocumentRequest

	baseURL := config.AppConfig.Webclient.DocumentService.BaseURL
	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.CreateDocumentRequest
	url := fmt.Sprintf("%v%v%v", baseURL, base, path)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if err := c.httpClient.Post(req, &response, url, headers); err != nil {
		fmt.Printf("Error in POST request: %v\n", err)
		return nil, err
	}

	// Return the first document request from the array
	if len(response) > 0 {
		return &response[0], nil
	}
	return nil, fmt.Errorf("no document requests were created")
}

func (c *DocumentRequestServiceClientImpl) DeleteDocumentRequest(id string) error {
	baseURL := config.AppConfig.Webclient.DocumentService.BaseURL
	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.DeleteDocumentRequest
	url := fmt.Sprintf("%v%v%v", baseURL, base, strings.Replace(path, "{id}", id, 1))
	headers := map[string]string{
		"Accept": "application/json",
	}

	if err := c.httpClient.Delete(nil, url, headers); err != nil {
		// If the error is about null response body, that's actually a success
		if strings.Contains(err.Error(), "unexpected end of JSON input") {
			return nil
		}
		fmt.Printf("Error in DELETE request: %v\n", err)
		return err
	}

	return nil
}

func (c *DocumentRequestServiceClientImpl) ListDocumentRequests(page, pageSize int) (*model.DocumentRequestListResponse, error) {
	var response model.DocumentRequestListResponse

	baseURL := config.AppConfig.Webclient.DocumentService.BaseURL
	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.ListDocumentRequests
	url := fmt.Sprintf("%v%v%v?page=%d&page_size=%d", baseURL, base, path, page, pageSize)
	headers := map[string]string{}

	if err := c.httpClient.Get(&response, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *DocumentRequestServiceClientImpl) ValidateDocumentRequest(id string, req model.ValidateDocumentRequestRequest) (*model.DocumentRequest, error) {
	var response model.DocumentRequest

	baseURL := config.AppConfig.Webclient.DocumentService.BaseURL
	base := config.AppConfig.Webclient.DocumentService.URL.Base
	path := config.AppConfig.Webclient.DocumentService.URL.ValidateDocumentRequest
	url := fmt.Sprintf("%v%v%v", baseURL, base, strings.Replace(path, "{id}", id, 1))
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if err := c.httpClient.Post(req, &response, url, headers); err != nil {
		fmt.Printf("Error in POST request: %v\n", err)
		return nil, err
	}

	return &response, nil
}
