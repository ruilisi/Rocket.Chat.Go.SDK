package rest

import (
	"net/url"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type SettingsMessageResponse struct {
	Status
	models.SettingsMessage
}

type SettingsOauthResponse struct {
	Status
	Services []models.SettingsOauth `json:"services"`
}

type SettingsResponse struct {
	base
	Settings []models.SettingsPrivate `json:"settings"`
}

type SettingsConfigurations struct {
	Configurations []interface{} `json:"configurations"`
	Status
}

type SettingsUpdateResonse struct {
	Status
}

func (c *Client) Settings(id string) (*SettingsMessageResponse, error) {
	response := new(SettingsMessageResponse)
	url := "settings/" + id
	if err := c.Get(url, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) PrivateSettings() (*SettingsResponse, error) {
	response := new(SettingsResponse)
	if err := c.Get("settings", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) PublicSettings(params url.Values) (*SettingsResponse, error) {
	response := new(SettingsResponse)
	if err := c.Get("settings.public", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) SettingsConfigurations() (*SettingsConfigurations, error) {
	response := new(SettingsConfigurations)
	if err := c.Get("service.configurations", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) SettingsUpdate(params url.Values) (*SettingsUpdateResonse, error) {
	response := new(SettingsUpdateResonse)
	url := "settings/" + params["_id"][0]
	if err := c.Get(url, params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) SettingsOauth() (*SettingsOauthResponse, error) {
	response := new(SettingsOauthResponse)
	if err := c.Get("settings.oauth", nil, response); err != nil {
		return nil, err
	}
	return response, nil

}
