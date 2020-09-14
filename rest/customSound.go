package rest

import "github.com/ruilisi/Rocket.Chat.Go.SDK/models"

type CustomSoundResponse struct {
	Sounds []models.Sound `json:"sounds"`
	base
}

func (c *Client) CustomSound() (*CustomSoundResponse, error) {
	response := new(CustomSoundResponse)
	if err := c.Get("custom-sounds.list", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
