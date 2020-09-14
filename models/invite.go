package models

import "time"

type CreateInviteRequest struct {
	Rid     string `json:"rid"`
	Day     int    `json:"days"` // 0 represent unlimited
	MaxUses int    `json:"maxUses"`
}

type Invite struct {
	ID        string    `json:"_id"`
	Days      int       `json:"days"`
	MaxUses   int       `json:"maxUses"`
	Rid       string    `json:"rid"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Expires   time.Time `json:"expires"`
	UpdatedAt time.Time `json:"_updatedAt"`
	Uses      int       `json:"uses"`
	URL       string    `json:"url"`
}
