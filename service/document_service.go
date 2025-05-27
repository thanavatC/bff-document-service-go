package service

import (
	"fmt"

	"github.com/thanavatC/bff-document-service-go/client"
	"github.com/thanavatC/bff-document-service-go/model"
)

type DocumentService interface {
	GetDocumentStatus(id string) (string, error)
	ReTranslateDocument(id string) (*model.Document, error)
	UpdateDocument(id string, req model.UpdateDocumentRequest) (*model.Document, error)
	DeleteDocument(id string) error
	ValidateDocument(id string, req model.ValidateDocumentRequest) (*model.Document, error)
}

type DocumentServiceImpl struct {
	documentServiceClient client.DocumentServiceClient
}

func NewDocumentServiceImpl(client client.DocumentServiceClient) DocumentService {
	return &DocumentServiceImpl{
		documentServiceClient: client,
	}
}

func (s *DocumentServiceImpl) GetDocumentStatus(id string) (string, error) {
	resp, err := s.documentServiceClient.GetDocumentStatus(id)
	if err != nil {
		fmt.Printf("Error calling GET document-service/document/%s/status: %v\n", id, err)
		return "", err
	}
	return resp, nil
}

func (s *DocumentServiceImpl) ReTranslateDocument(id string) (*model.Document, error) {
	resp, err := s.documentServiceClient.ReTranslateDocument(id)
	if err != nil {
		fmt.Printf("Error calling POST document-service/document/%s/retranslate: %v\n", id, err)
		return nil, err
	}
	return resp, nil
}

func (s *DocumentServiceImpl) UpdateDocument(id string, req model.UpdateDocumentRequest) (*model.Document, error) {
	resp, err := s.documentServiceClient.UpdateDocument(id, req)
	if err != nil {
		fmt.Printf("Error calling PATCH document-service/document/%s: %v\n", id, err)
		return nil, err
	}
	return resp, nil
}

func (s *DocumentServiceImpl) DeleteDocument(id string) error {
	err := s.documentServiceClient.DeleteDocument(id)
	if err != nil {
		fmt.Printf("Error calling DELETE document-service/document/%s: %v\n", id, err)
		return err
	}
	return nil
}

func (s *DocumentServiceImpl) ValidateDocument(id string, req model.ValidateDocumentRequest) (*model.Document, error) {
	resp, err := s.documentServiceClient.ValidateDocument(id, req)
	if err != nil {
		fmt.Printf("Error calling POST document-service/document/%s/validate: %v\n", id, err)
		return nil, err
	}
	return resp, nil
}
