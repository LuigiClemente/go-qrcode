package models

import "time"

type QRCode struct {
	ID        int       `json:"id"`
	Token     string    `json:"token"`
	Image     string    `json:"image"`
	Url       string    `json:"url"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
}
