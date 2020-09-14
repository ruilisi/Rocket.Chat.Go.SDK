package models

import "time"

type Sound struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	UpdatedAt time.Time `json:"_updatedAt"`
}
