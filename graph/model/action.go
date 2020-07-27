package model

import "time"

type Action struct {
	ID              string    `json:"id"`
	ActionTimestamp time.Time `json:"actionTimestamp"`
	ActivityID      string    `json:"activity"`
}
