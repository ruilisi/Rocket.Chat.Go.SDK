package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type EmojiUpdateResponse struct {
	Status
}

type EmojiDeleteResponse struct {
	Status
}

type EmojiListResponse struct {
	Status
	Emojis struct {
		Update []struct {
			ID        string        `json:"_id"`
			Name      string        `json:"name"`
			Aliases   []interface{} `json:"aliases"`
			Extension string        `json:"extension"`
			UpdatedAt time.Time     `json:"_updatedAt"`
		} `json:"update"`
		Remove []interface{} `json:"remove"`
	} `json:"emojis"`
}

type EmojiCreateResponse struct {
	Status
}

func (c *Client) EmojiCreate(params map[string]io.Reader) (*EmojiCreateResponse, error) {
	response := new(EmojiCreateResponse)
	if err := c.PostFormData(http.MethodPost, "emoji-custom.create", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) EmojiList(params url.Values) (*EmojiListResponse, error) {

	response := new(EmojiListResponse)
	if err := c.Get("emoji-custom.list", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) EmojiDelete(req *models.EmojiDeleteRequest) (*EmojiDeleteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(EmojiDeleteResponse)
	err = c.Post("emoji-custom.delete", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) EmojiUpdate(params map[string]io.Reader) (*EmojiUpdateResponse, error) {
	response := new(EmojiUpdateResponse)
	if err := c.PostFormData(http.MethodPost, "emoji-custom.update", params, response); err != nil {
		return nil, err
	}
	return response, nil
}
