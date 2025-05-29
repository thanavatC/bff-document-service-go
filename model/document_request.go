package model

import (
	"time"

	"gorm.io/gorm"
)

type DocumentRequestStatus string

const (
	DocumentRequestStatusPending  DocumentRequestStatus = "pending"
	DocumentRequestStatusApproved DocumentRequestStatus = "approved"
	DocumentRequestStatusRejected DocumentRequestStatus = "rejected"
)

type DocumentRequest struct {
	ID           string                `gorm:"primaryKey" json:"id"`
	CompanyID    string                `gorm:"index" json:"company_id"`
	CompanyName  string                `json:"company_name"`
	DocumentID   string                `gorm:"index" json:"document_id"`
	Name         string                `json:"name"`
	Type         string                `json:"type"`
	FileName     string                `json:"file_name"`
	DocumentType string                `json:"document_type"`
	Status       DocumentRequestStatus `json:"status"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	DeletedAt    gorm.DeletedAt        `gorm:"index" json:"-"`
}

func (DocumentRequest) TableName() string {
	return "document_requests"
}

type CreateDocumentRequestRequest struct {
	CompanyID     string   `json:"companyId" binding:"required"`
	CompanyName   string   `json:"companyName" binding:"required"`
	Files         []string `json:"files" binding:"required"`
	DocumentTypes []string `json:"documentTypes" binding:"required"`
}

type UpdateDocumentRequestRequest struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

type ValidateDocumentRequestRequest struct {
	IsValid bool `json:"is_valid" binding:"required"`
}

type DocumentRequestResponse struct {
	ID           string                `json:"id"`
	CompanyID    string                `json:"company_id"`
	CompanyName  string                `json:"company_name"`
	DocumentID   string                `json:"document_id"`
	FileName     string                `json:"file_name"`
	DocumentType string                `json:"document_type"`
	Status       DocumentRequestStatus `json:"status"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
}

type DocumentRequestListResponse struct {
	Requests []DocumentRequestResponse `json:"requests"`
	Total    int64                     `json:"total"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"page_size"`
}

type ValidateDocumentRequestResponse struct {
	ID         string    `json:"id"`
	CompanyID  string    `json:"company_id"`
	DocumentID string    `json:"document_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
