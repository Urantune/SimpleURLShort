package models

import "time"

type Link struct {
	ID          int       `json:"-"`
	Code        string    `json:"code"`
	OriginalURL string    `json:"original_url"`
	Visits      int       `json:"visits"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
