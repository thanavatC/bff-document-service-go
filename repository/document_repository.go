package repository

import (
	"fmt"

	"github.com/thanavatC/bff-document-service-go/model"
	"gorm.io/gorm"
)

type DocumentRepository interface {
	SaveDocument(doc *model.Document) error
	GetDocument(id string) (*model.Document, error)
	ListDocuments(offset, limit int) ([]model.Document, int64, error)
	GetDocumentByID(id string) (*model.Document, error)
	DocumentExists(id string) (bool, error)
	DeleteDocument(id string) error
	UpdateDocument(doc *model.Document) (*model.Document, error)
}

type DocumentRepositoryImpl struct {
	db *gorm.DB
}

func NewDocumentRepositoryImpl(db *gorm.DB) DocumentRepository {
	return &DocumentRepositoryImpl{
		db: db,
	}
}

func (repo *DocumentRepositoryImpl) SaveDocument(doc *model.Document) error {
	return repo.db.Create(doc).Error
}

func (repo *DocumentRepositoryImpl) GetDocument(id string) (*model.Document, error) {
	var doc model.Document
	if err := repo.db.First(&doc, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("document not found: %v", err)
	}
	return &doc, nil
}

func (repo *DocumentRepositoryImpl) ListDocuments(offset, limit int) ([]model.Document, int64, error) {
	var docs []model.Document
	var total int64

	// Get total count
	if err := repo.db.Model(&model.Document{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	if err := repo.db.Offset(offset).Limit(limit).Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return docs, total, nil
}

func (repo *DocumentRepositoryImpl) GetDocumentByID(id string) (*model.Document, error) {
	var doc model.Document
	result := repo.db.Unscoped().Where("id = ?", id).First(&doc)
	if result.Error != nil {
		return nil, result.Error
	}
	return &doc, nil
}

func (repo *DocumentRepositoryImpl) DocumentExists(id string) (bool, error) {
	var count int64
	result := repo.db.Model(&model.Document{}).Where("id = ?", id).Count(&count)
	return count > 0, result.Error
}

func (repo *DocumentRepositoryImpl) DeleteDocument(id string) error {
	result := repo.db.Where("id = ?", id).Delete(&model.Document{})
	return result.Error
}

func (repo *DocumentRepositoryImpl) UpdateDocument(doc *model.Document) (*model.Document, error) {
	result := repo.db.Save(doc)
	if result.Error != nil {
		return nil, result.Error
	}
	return doc, nil
}
