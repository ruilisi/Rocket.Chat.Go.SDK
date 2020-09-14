package rest

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type CreateIntegrationResponse struct {
	Status
	Integration struct {
		Type          string        `json:"type"`
		Name          string        `json:"name"`
		Enabled       bool          `json:"enabled"`
		Username      string        `json:"username"`
		Event         string        `json:"event"`
		Urls          []string      `json:"urls"`
		ScriptEnabled bool          `json:"scriptEnabled"`
		UserID        string        `json:"userId"`
		Channel       []interface{} `json:"channel"`
		CreatedAt     time.Time     `json:"_createdAt"`
		CreatedBy     struct {
			Username string `json:"username"`
			ID       string `json:"_id"`
		} `json:"_createdBy"`
		UpdatedAt time.Time `json:"_updatedAt"`
		ID        string    `json:"_id"`
	} `json:"integration"`
}

type GetIntegrationRequest struct {
	ID       string `json:"integrationId"`
	CreateBy string `json:"createBy"`
}

type GetIntegrationResponse struct {
	Integration struct {
		ID            string    `json:"_id"`
		Type          string    `json:"type"`
		Name          string    `json:"name"`
		Event         string    `json:"event"`
		Enabled       bool      `json:"enabled"`
		Username      string    `json:"username"`
		Urls          []string  `json:"urls"`
		ScriptEnabled bool      `json:"scriptEnabled"`
		Channel       []string  `json:"channel"`
		UserID        string    `json:"userId"`
		CreatedAt     time.Time `json:"_createdAt"`
		CreatedBy     struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"_createdBy"`
		UpdatedAt time.Time `json:"_updatedAt"`
	} `json:"integration"`
	Status
}

type HistoryIntegrationResponse struct {
	base
	History []models.IntegrationHistory `json:"history"`
}

type ListIntegrationResponse struct {
	base
	Integration []models.Integration `json:"integrations"`
}

type RemoveIntegrationRequset struct {
	Type        string `json:"type"`
	Integration string `json:"integrationId"`
}

type RemoveIntegrationResponse struct {
	Status
	Integration models.Integration `json:"integration"`
}

func (c *Client) CreateIntegration(req *models.CreateIntegration) (*CreateIntegrationResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(CreateIntegrationResponse)
	err = c.Post("integrations.create", bytes.NewBuffer(body), response)
	return response, err

}

func (c *Client) GetIntegration(params url.Values) (*GetIntegrationResponse, error) {

	response := new(GetIntegrationResponse)
	if err := c.Get("integrations.get", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) HistoryIntegration(params url.Values) (*HistoryIntegrationResponse, error) {

	response := new(HistoryIntegrationResponse)
	if err := c.Get("integrations.history", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ListIntegration() (*ListIntegrationResponse, error) {
	response := new(ListIntegrationResponse)
	if err := c.Get("integrations.list", nil, response); err != nil {
		return nil, err
	}

	return response, nil

}

func (c *Client) RemoveIntegration(req *RemoveIntegrationRequset) (*RemoveIntegrationResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(RemoveIntegrationResponse)
	err = c.Post("integrations.remove", bytes.NewBuffer(body), response)
	return response, err

}
