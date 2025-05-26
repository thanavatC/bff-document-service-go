package repository

import (
	"github.com/SPVJ/document-service-go/model"
	"gorm.io/gorm"
)

type DocumentRequestRepository interface {
	CreateDocumentRequest(docRequest *model.DocumentRequest) (*model.DocumentRequest, error)
	ListDocumentRequests(page, pageSize int) ([]model.DocumentRequest, int64, error)
	DocumentRequestExists(id string) (bool, error)
	DeleteDocumentRequest(id string) error
	GetDocumentRequestByID(id string) (*model.DocumentRequest, error)
	UpdateDocumentRequest(docRequest *model.DocumentRequest) (*model.DocumentRequest, error)
}

type DocumentRequestRepositoryImpl struct {
	db *gorm.DB
}

func NewDocumentRequestRepositoryImpl(db *gorm.DB) DocumentRequestRepository {
	return &DocumentRequestRepositoryImpl{
		db: db,
	}
}

func (repo *DocumentRequestRepositoryImpl) CreateDocumentRequest(docRequest *model.DocumentRequest) (*model.DocumentRequest, error) {
	result := repo.db.Create(docRequest)
	if result.Error != nil {
		return nil, result.Error
	}
	return docRequest, nil
}

func (repo *DocumentRequestRepositoryImpl) ListDocumentRequests(page, pageSize int) ([]model.DocumentRequest, int64, error) {
	var requests []model.DocumentRequest
	var total int64

	offset := (page - 1) * pageSize

	repo.db.Model(&model.DocumentRequest{}).Count(&total)
	result := repo.db.Offset(offset).Limit(pageSize).Find(&requests)

	return requests, total, result.Error
}

func (repo *DocumentRequestRepositoryImpl) DocumentRequestExists(id string) (bool, error) {
	var count int64
	result := repo.db.Unscoped().Model(&model.DocumentRequest{}).Where("id = ?", id).Count(&count)
	return count > 0, result.Error
}

func (repo *DocumentRequestRepositoryImpl) DeleteDocumentRequest(id string) error {
	var docRequest model.DocumentRequest
	result := repo.db.Unscoped().Where("id = ?", id).First(&docRequest)
	if result.Error != nil {
		return result.Error
	}
	return repo.db.Delete(&docRequest).Error
}

func (repo *DocumentRequestRepositoryImpl) GetDocumentRequestByID(id string) (*model.DocumentRequest, error) {
	var docRequest model.DocumentRequest
	result := repo.db.Unscoped().Where("id = ?", id).First(&docRequest)
	if result.Error != nil {
		return nil, result.Error
	}
	return &docRequest, nil
}

func (repo *DocumentRequestRepositoryImpl) UpdateDocumentRequest(docRequest *model.DocumentRequest) (*model.DocumentRequest, error) {
	result := repo.db.Save(docRequest)
	if result.Error != nil {
		return nil, result.Error
	}
	return docRequest, nil
}
