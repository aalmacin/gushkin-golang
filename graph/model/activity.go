package model

type Activity struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Positive    bool   `json:"positive"`
	FundAmt     int    `json:"fundAmt"`
}
