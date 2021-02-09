package model

// Product model
type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
}

// Products is a list of the object product
type Products struct {
	Products []Product `json:"products"`
}
