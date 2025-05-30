package service

import (
	"fmt"

	"github.com/thanavatC/bff-document-service-go/client"
	"github.com/thanavatC/bff-document-service-go/model"
)

type DocumentRequestService interface {
	CreateDocumentRequest(req model.CreateDocumentRequestRequest) (*model.DocumentRequest, error)
	DeleteDocumentRequest(id string) error
	ListDocumentRequests(page, pageSize int) ([]*model.DocumentRequest, int64, error)
	ValidateDocumentRequest(id string, req model.ValidateDocumentRequestRequest) (*model.DocumentRequest, error)
}

type DocumentRequestServiceImpl struct {
	documentRequestServiceClient client.DocumentRequestServiceClient
}

func NewDocumentRequestServiceImpl(client client.DocumentRequestServiceClient) DocumentRequestService {
	return &DocumentRequestServiceImpl{
		documentRequestServiceClient: client,
	}
}

func (s *DocumentRequestServiceImpl) CreateDocumentRequest(req model.CreateDocumentRequestRequest) (*model.DocumentRequest, error) {
	resp, err := s.documentRequestServiceClient.CreateDocumentRequest(req)
	if err != nil {
		fmt.Printf("Error calling POST document-request-service/document-request: %v\n", err)
		return nil, err
	}
	return resp, nil
}

func (s *DocumentRequestServiceImpl) DeleteDocumentRequest(id string) error {
	err := s.documentRequestServiceClient.DeleteDocumentRequest(id)
	if err != nil {
		fmt.Printf("Error calling DELETE document-request-service/document-request/%s: %v\n", id, err)
		return err
	}
	return nil
}

func (s *DocumentRequestServiceImpl) ListDocumentRequests(page, pageSize int) ([]*model.DocumentRequest, int64, error) {
	resp, err := s.documentRequestServiceClient.ListDocumentRequests(page, pageSize)
	if err != nil {
		fmt.Printf("Error calling GET document-request-service/document-requests: %v\n", err)
		return nil, 0, err
	}

	// Convert DocumentRequestResponse to DocumentRequest
	requests := make([]*model.DocumentRequest, len(resp.Requests))
	for i, req := range resp.Requests {
		requests[i] = &model.DocumentRequest{
			ID:           req.ID,
			CompanyID:    req.CompanyID,
			CompanyName:  req.CompanyName,
			DocumentID:   req.DocumentID,
			FileName:     req.FileName,
			DocumentType: req.DocumentType,
			Status:       req.Status,
			CreatedAt:    req.CreatedAt,
			UpdatedAt:    req.UpdatedAt,
		}
	}
	return requests, resp.Total, nil
}

func (s *DocumentRequestServiceImpl) ValidateDocumentRequest(id string, req model.ValidateDocumentRequestRequest) (*model.DocumentRequest, error) {
	resp, err := s.documentRequestServiceClient.ValidateDocumentRequest(id, req)
	if err != nil {
		fmt.Printf("Error calling POST document-request-service/document-request/%s/validate: %v\n", id, err)
		return nil, err
	}
	return resp, nil
}
