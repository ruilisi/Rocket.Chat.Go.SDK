package models

import "time"

type OauthApp struct {
	ID           string    `json:"_id"`
	Name         string    `json:"name"`
	Active       bool      `json:"active"`
	ClientID     string    `json:"clientId"`
	ClientSecret string    `json:"clientSecret"`
	RedirectURI  string    `json:"redirectUri"`
	CreatedAt    time.Time `json:"_createdAt"`
	CreatedBy    struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
	} `json:"_createdBy"`
	UpdatedAt time.Time `json:"_updatedAt"`
}
