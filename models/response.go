package models

type SuccessResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"Operation completed successfully"`
	Data    interface{} `json:"data"`
}

type SuccessResponseTest struct {
	Status  bool        `json:"status" example:"true"`
	Message string      `json:"message" example:"Operation completed successfully"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"An error occurred"`
}
