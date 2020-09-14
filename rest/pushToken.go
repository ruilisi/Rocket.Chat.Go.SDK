package rest

import (
	"bytes"
	"encoding/json"
	"time"
)

type PushTokenPostRequest struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Value   string `json:"value"`
	AppName string `json:"appName"`
}

type PushTokenPostResponse struct {
	Result struct {
		Token struct {
			Gcm string `json:"gcm"`
		} `json:"token"`
		AppName   string    `json:"appName"`
		UserID    string    `json:"userId"`
		Enabled   bool      `json:"enabled"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		ID        string    `json:"_id"`
	} `json:"result"`
	Status
}
type PushTokenDeleteRequest struct {
	Token string `json:"token"`
}

type PushTokenDeleteResponse struct {
	Status
}

func (c *Client) PushTokenPost(req *PushTokenPostRequest) (*PushTokenPostResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(PushTokenPostResponse)
	err = c.Post("push.token", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) PushTokenDelete(req *PushTokenDeleteRequest) (*PushTokenDeleteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(PushTokenDeleteResponse)
	if err := c.Delete("push.token", bytes.NewBuffer(body), response); err != nil {
		return nil, err
	}
	return response, nil

}
