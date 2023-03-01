package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListResponse struct {
	TotalPage  interface{} `json:"total_page"`
	TotalData  interface{} `json:"total_data"`
	DataShowed interface{} `json:"data_showed"`
	PerPage    interface{} `json:"per_page"`
	Page       interface{} `json:"page"`
	Records    interface{} `json:"records"`
}
