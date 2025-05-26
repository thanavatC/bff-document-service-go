-- Create companies table
CREATE TABLE IF NOT EXISTS companies (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create documents table
CREATE TABLE IF NOT EXISTS documents (
    id VARCHAR(36) PRIMARY KEY,
    company_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    file_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

-- Create document_requests table
CREATE TABLE IF NOT EXISTS document_requests (
    id VARCHAR(36) PRIMARY KEY,
    company_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

-- Insert test companies
INSERT INTO companies (id, name) VALUES
    ('comp_001', 'Test Company 1'),
    ('comp_002', 'Test Company 2'),
    ('comp_003', 'Test Company 3');

-- Insert test documents
INSERT INTO documents (id, company_id, name, type, status, file_url) VALUES
    ('doc_001', 'comp_001', 'Test Document 1', 'pdf', 'pending', 'https://example.com/doc1.pdf'),
    ('doc_002', 'comp_001', 'Test Document 2', 'docx', 'approved', 'https://example.com/doc2.docx'),
    ('doc_003', 'comp_002', 'Test Document 3', 'pdf', 'rejected', 'https://example.com/doc3.pdf');

-- Insert test document requests
INSERT INTO document_requests (id, company_id, name, type, status) VALUES
    ('req_001', 'comp_001', 'Test Request 1', 'pdf', 'pending'),
    ('req_002', 'comp_002', 'Test Request 2', 'docx', 'approved'),
    ('req_003', 'comp_003', 'Test Request 3', 'pdf', 'rejected'); 