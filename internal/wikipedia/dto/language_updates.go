package dto

import "time"

type LanguageUpdate struct {
	Language  string    `json:"language"`
	UpdatedAt time.Time `json:"updated_at"`
}
