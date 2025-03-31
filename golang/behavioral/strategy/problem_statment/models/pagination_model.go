package models

type Pagination struct {
	Limit    int `json:"limit"`
	PageSize int `json:"pageSize"`
}
