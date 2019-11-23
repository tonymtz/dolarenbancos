package models

import "time"

type Price struct {
	Id        int64     `json:"id"`
	Sell      float64   `json:"sell"`
	Buy       float64   `json:"buy"`
	CreatedAt time.Time `json:"created_at"`
	Bank      int       `json:"bank"`
}
