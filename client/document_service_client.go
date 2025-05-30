package client

import (
	"fmt"
	"strings"

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

func (c *DocumentServiceClientImpl) constructURL(path string, id string) string {
	baseURL := config.AppConfig.Webclient.DocumentService.BaseURL
	base := config.AppConfig.Webclient.DocumentService.URL.Base
	return fmt.Sprintf("%v%v%v", baseURL, base, strings.Replace(path, "{id}", id, 1))
}

func (c *DocumentServiceClientImpl) GetDocumentStatus(id string) (string, error) {
	var response struct {
		Status string `json:"status"`
	}

	url := c.constructURL(config.AppConfig.Webclient.DocumentService.URL.GetDocumentStatus, id)
	headers := map[string]string{}

	if err := c.httpClient.Get(&response, url, headers); err != nil {
		return "", err
	}

	return response.Status, nil
}

func (c *DocumentServiceClientImpl) ReTranslateDocument(id string) (*model.Document, error) {
	var response model.Document

	url := c.constructURL(config.AppConfig.Webclient.DocumentService.URL.ReTranslateDocument, id)
	headers := map[string]string{}

	if err := c.httpClient.Post(&response, nil, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *DocumentServiceClientImpl) UpdateDocument(id string, req model.UpdateDocumentRequest) (*model.Document, error) {
	var response model.Document

	url := c.constructURL(config.AppConfig.Webclient.DocumentService.URL.UpdateDocument, id)
	headers := map[string]string{}

	if err := c.httpClient.Put(req, &response, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *DocumentServiceClientImpl) DeleteDocument(id string) error {
	url := c.constructURL(config.AppConfig.Webclient.DocumentService.URL.DeleteDocument, id)
	headers := map[string]string{}

	// If the error is about null response body (204 No Content), that's actually a success
	if err := c.httpClient.Delete(nil, url, headers); err != nil {
		if strings.Contains(err.Error(), "unexpected end of JSON input") {
			return nil
		}
		return err
	}

	return nil
}

func (c *DocumentServiceClientImpl) ValidateDocument(id string, req model.ValidateDocumentRequest) (*model.Document, error) {
	var response model.Document

	url := c.constructURL(config.AppConfig.Webclient.DocumentService.URL.ValidateDocument, id)
	headers := map[string]string{}

	if err := c.httpClient.Post(req, &response, url, headers); err != nil {
		return nil, err
	}

	return &response, nil
}
