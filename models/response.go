package models

type Response struct {
	Page         int        `json:"page"`
	PageSize     int        `json:"pageSize"`
	TotalResults int        `json:"totalResults"`
	TotalPages   int        `json:"totalPages"`
	Properties   []Property `json:"properties"`
}
