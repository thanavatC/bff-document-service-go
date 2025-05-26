package service

import (
	"fmt"
	"time"

	"github.com/SPVJ/document-service-go/model"
	"github.com/SPVJ/document-service-go/repository"
	"github.com/google/uuid"
)

type DocumentRequestService struct {
	repo repository.DocumentRequestRepository
}

func NewDocumentRequestService(repo repository.DocumentRequestRepository) *DocumentRequestService {
	return &DocumentRequestService{
		repo: repo,
	}
}

func (s *DocumentRequestService) CreateDocumentRequest(req model.CreateDocumentRequestRequest) (*model.DocumentRequest, error) {
	// Generate unique ID
	id := uuid.New().String()

	// Create document request
	docRequest := &model.DocumentRequest{
		ID:         id,
		Name:       req.Name,
		Type:       req.Type,
		CompanyID:  req.CompanyID,
		DocumentID: req.DocumentID,
		Status:     model.DocumentRequestStatusPending,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return s.repo.CreateDocumentRequest(docRequest)
}

func (s *DocumentRequestService) ListDocumentRequests(page, pageSize int) ([]model.DocumentRequest, int64, error) {
	return s.repo.ListDocumentRequests(page, pageSize)
}

func (s *DocumentRequestService) DeleteDocumentRequest(id string) error {
	exists, err := s.repo.DocumentRequestExists(id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("document request not found")
	}
	return s.repo.DeleteDocumentRequest(id)
}

func (s *DocumentRequestService) ValidateDocumentRequest(id string, req model.ValidateDocumentRequestRequest) (*model.DocumentRequest, error) {
	docRequest, err := s.repo.GetDocumentRequestByID(id)
	if err != nil {
		return nil, err
	}

	status := model.DocumentRequestStatusRejected
	if req.IsValid {
		status = model.DocumentRequestStatusApproved
	}

	docRequest.Status = status
	docRequest.UpdatedAt = time.Now()

	return s.repo.UpdateDocumentRequest(docRequest)
}
