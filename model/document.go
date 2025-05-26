package model

import (
	"time"

	"gorm.io/gorm"
)

type DocumentStatus string

const (
	DocumentStatusPending     DocumentStatus = "pending"
	DocumentStatusTranslating DocumentStatus = "translating"
	DocumentStatusCompleted   DocumentStatus = "completed"
	DocumentStatusFailed      DocumentStatus = "failed"
	DocumentStatusRejected    DocumentStatus = "rejected"
	DocumentStatusApproved    DocumentStatus = "approved"
)

type Document struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	CompanyID string         `gorm:"index" json:"company_id"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	FileURL   string         `json:"file_url"`
	Status    DocumentStatus `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Document) TableName() string {
	return "document"
}

type UploadDocumentRequest struct {
	Name    string `json:"name" binding:"required"`
	Type    string `json:"type" binding:"required"`
	FileURL string `json:"file_url" binding:"required"`
}

type DocumentResponse struct {
	ID        string         `json:"id"`
	CompanyID string         `json:"company_id"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	FileURL   string         `json:"file_url"`
	Status    DocumentStatus `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type DocumentListResponse struct {
	Documents []DocumentResponse `json:"documents"`
}

type TranslateDocumentRequest struct {
	TargetLanguage string `json:"target_language" binding:"required"`
}

type DocumentStatusResponse struct {
	ID     string         `json:"id"`
	Status DocumentStatus `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UpdateDocumentRequest struct {
	Name   string         `json:"name"`
	Type   string         `json:"type"`
	Status DocumentStatus `json:"status"`
}

type ValidateDocumentRequest struct {
	IsValid bool `json:"is_valid" binding:"required"`
}

type ValidateDocumentResponse struct {
	ID        string         `json:"id"`
	CompanyID string         `json:"company_id"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	FileURL   string         `json:"file_url"`
	Status    DocumentStatus `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
