package model

// SMS model
type SMS struct {
	Message     string `json:"message"`
	PhoneNumber string `json:"phoneNumber"`
}
