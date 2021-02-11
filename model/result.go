package model

// Result model
type Result struct {
	Code    int         `json:"code" `
	Error   bool        `json:"error" example:"false"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
