package model

type Client struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	City  string `json:"city"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
