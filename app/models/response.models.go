package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListResponse struct {
	TotalPage    interface{} `json:"total_page"`
	TotalData    interface{} `json:"total_data"`
	TotalRecords interface{} `json:"total_records"`
	PerPage      interface{} `json:"per_page"`
	Page         interface{} `json:"page"`
	Records      interface{} `json:"records"`
}
