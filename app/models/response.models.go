package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListResponse struct {
	TotalPage interface{} `json:"total_page"`
	TotalData interface{} `json:"total_data"`
	PerPage   interface{} `json:"per_page"`
	Page      interface{} `json:"page"`
	Record    interface{} `json:"record"`
}
