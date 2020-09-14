package rest

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type GetSubscriptionsResponse struct {
	Status
	Update []struct {
		ID            string    `json:"_id"`
		Open          bool      `json:"open"`
		Alert         bool      `json:"alert"`
		Unread        int       `json:"unread"`
		UserMentions  int       `json:"userMentions"`
		GroupMentions int       `json:"groupMentions"`
		Ts            time.Time `json:"ts,omitempty"`
		Rid           string    `json:"rid"`
		Name          string    `json:"name"`
		Fname         string    `json:"fname"`
		T             string    `json:"t"`
		U             struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Ls                      time.Time     `json:"ls,omitempty"`
		UpdatedAt               time.Time     `json:"_updatedAt"`
		Roles                   []string      `json:"roles,omitempty"`
		DesktopNotifications    string        `json:"desktopNotifications"`
		MobilePushNotifications string        `json:"mobilePushNotifications"`
		EmailNotifications      string        `json:"emailNotifications"`
		Prid                    string        `json:"prid,omitempty"`
		F                       bool          `json:"f,omitempty"`
		Ignored                 []interface{} `json:"ignored,omitempty"`
	} `json:"update"`
	Remove []interface{} `json:"remove"`
}

type GetOneResponse struct {
	Status
	Subscription struct {
		ID            string    `json:"_id"`
		Open          bool      `json:"open"`
		Alert         bool      `json:"alert"`
		Unread        int       `json:"unread"`
		UserMentions  int       `json:"userMentions"`
		GroupMentions int       `json:"groupMentions"`
		Ts            time.Time `json:"ts"`
		Rid           string    `json:"rid"`
		Name          string    `json:"name"`
		Fname         string    `json:"fname"`
		CustomFields  struct {
		} `json:"customFields"`
		T string `json:"t"`
		U struct {
			ID       string      `json:"_id"`
			Username string      `json:"username"`
			Name     interface{} `json:"name"`
		} `json:"u"`
		Ls                          time.Time `json:"ls"`
		UpdatedAt                   time.Time `json:"_updatedAt"`
		Roles                       []string  `json:"roles"`
		DisableNotifications        bool      `json:"disableNotifications"`
		EmailNotifications          string    `json:"emailNotifications"`
		AudioNotificationValue      string    `json:"audioNotificationValue"`
		DesktopNotificationDuration int       `json:"desktopNotificationDuration"`
		AudioNotifications          string    `json:"audioNotifications"`
		MobilePushNotifications     string    `json:"mobilePushNotifications"`
		F                           bool      `json:"f"`
		Meta                        struct {
			Revision int   `json:"revision"`
			Created  int64 `json:"created"`
			Version  int   `json:"version"`
		} `json:"meta"`
	} `json:"subscription"`
}

type ReadResp struct {
	Status
}

type UnReadResp struct {
	Status
}

func (c *Client) GetSubscriptions(params url.Values) (*GetSubscriptionsResponse, error) {
	response := new(GetSubscriptionsResponse)
	if err := c.Get("subscriptions.get", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetOne(params url.Values) (*GetOneResponse, error) {
	response := new(GetOneResponse)
	if err := c.Get("subscriptions.getOne", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Read(req *models.Read) (*ReadResp, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(ReadResp)
	err = c.Post("subscriptions.read", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) UnRead(req *models.UnRead) (*UnReadResp, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(UnReadResp)
	err = c.Post("subscriptions.unread", bytes.NewBuffer(body), response)
	return response, err
}
