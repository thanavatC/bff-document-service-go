package model

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (Company) TableName() string {
	return "company"
}

type CompanyListResponse struct {
	Companies []Company `json:"companies"`
}
