package rest

import (
	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type CustomUserStatusResponse struct {
	Status []models.CustomUserStatus `json:"statuses"`
	base
}

func (c *Client) CustomUserStatus() (*CustomUserStatusResponse, error) {
	response := new(CustomUserStatusResponse)
	if err := c.Get("custom-user-status.list", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
