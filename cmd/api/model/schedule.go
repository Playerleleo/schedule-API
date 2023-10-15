package model

import "time"

type schedule struct {
	ID     string  `json:"id"`
	Description string `json:"description"`
	DateTime time.Time `json:"date_time"`
}