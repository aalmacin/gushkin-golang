package model

type Wish struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Price       int     `json:"price"`
	Source      *string `json:"source"`
	Priority    string  `json:"priority"`
	Status      string  `json:"status"`
	UserID      string
}
