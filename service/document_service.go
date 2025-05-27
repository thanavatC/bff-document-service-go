package service

import (
	"fmt"
	"time"

	"github.com/thanavatC/bff-document-service-go/model"
	"github.com/thanavatC/bff-document-service-go/repository"
)

type DocumentService struct {
	repo repository.DocumentRepository
}

func NewDocumentService(repo repository.DocumentRepository) *DocumentService {
	return &DocumentService{
		repo: repo,
	}
}

func (s *DocumentService) GetDocumentStatus(id string) (model.DocumentStatus, error) {
	doc, err := s.repo.GetDocumentByID(id)
	if err != nil {
		return "", err
	}
	return doc.Status, nil
}

func (s *DocumentService) ReTranslateDocument(id string) (*model.Document, error) {
	doc, err := s.repo.GetDocumentByID(id)
	if err != nil {
		return nil, err
	}

	// Update status to indicate re-translation
	doc.Status = model.DocumentStatusTranslating
	doc.UpdatedAt = time.Now()

	return s.repo.UpdateDocument(doc)
}

func (s *DocumentService) UpdateDocument(id string, req model.UpdateDocumentRequest) (*model.Document, error) {
	doc, err := s.repo.GetDocumentByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		doc.Name = req.Name
	}
	if req.Type != "" {
		doc.Type = req.Type
	}
	if req.Status != "" {
		doc.Status = req.Status
	}
	doc.UpdatedAt = time.Now()

	return s.repo.UpdateDocument(doc)
}

func (s *DocumentService) DeleteDocument(id string) error {
	exists, err := s.repo.DocumentExists(id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("document not found")
	}
	return s.repo.DeleteDocument(id)
}

func (s *DocumentService) ValidateDocument(id string, req model.ValidateDocumentRequest) (*model.Document, error) {
	doc, err := s.repo.GetDocumentByID(id)
	if err != nil {
		return nil, err
	}

	status := model.DocumentStatusRejected
	if req.IsValid {
		status = model.DocumentStatusApproved
	}

	doc.Status = status
	doc.UpdatedAt = time.Now()

	return s.repo.UpdateDocument(doc)
}
