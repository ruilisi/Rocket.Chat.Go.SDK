package models

import "time"

type CustomUserStatus struct {
	ID         string    `json:"_id"`
	Name       string    `json:"name"`
	StatusType string    `json:"statusType"`
	UpdatedAt  time.Time `json:"_updatedAt"`
}
