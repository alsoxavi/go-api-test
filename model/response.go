package model

// ErrorModel model
type ErrorModel struct {
	Error   bool   `json:"error" example:"true"`
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
