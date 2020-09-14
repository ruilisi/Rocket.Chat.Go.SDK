package rest

import (
	"net/url"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type OauthAppGetResponse struct {
	OauthApp models.OauthApp `json:"oauthApp"`
	Status
}

type OauthAppListResponse struct {
	OauthApps []models.OauthApp `json:"oauthApps"`
	Status
}

func (c *Client) GetOauthApp(params url.Values) (*OauthAppGetResponse, error) {
	response := new(OauthAppGetResponse)
	if err := c.Get("oauth-apps.get", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) ListOauthApp() (*OauthAppListResponse, error) {
	response := new(OauthAppListResponse)
	if err := c.Get("oauth-apps.list", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
