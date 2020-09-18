package rest

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type AuthLoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Resume   string `jsonL:resume`
}

type AuthLoginResponse struct {
	Data models.AuthData `json:"data"`
	Status
}

type AuthLogoutResponse struct {
	Status
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

type AuthFacebookLoginRequest struct {
	ServiceName string `json:"serviceName"`
	AccessToken string `json:"accessToken"`
	Secret      string `json:"secret"`
	ExpiresIn   int    `json:"expiresIn"`
}

type AuthGoogleLoginRequest struct {
	ServiceName string `json:"serviceName"`
	AccessToken string `json:"accessToken"`
	IDToken     string `json:"idToken"`
	ExpiresIn   int    `json:"expiresIn"`
	Scope       string `json:"scope"`
}

type AuthTwitterLoginRequest struct {
	ServiceName       string `json:"serviceNae"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	AppSecret         string `json:"appSecret"`
	AppId             string `json:"appId"`
	ExpiresIn         int    `json:"expiresIn"`
}

type AuthMeResponse struct {
	models.Me
	Status
}

func (c *Client) AuthLogin(req AuthLoginRequest) (*AuthLoginResponse, error) {
	var data url.Values
	if req.Resume == "" {
		data = url.Values{"user": {req.User}, "password": {req.Password}}
	} else {
		data = url.Values{"resume": {req.Resume}}
	}

	response := new(AuthLoginResponse)
	c.PostForm("login", data, response)
	c.Auth = &AuthInfo{ID: response.Data.UserID, Token: response.Data.AuthToken}
	return response, nil
}

func (c *Client) AuthLogout() (*AuthLogoutResponse, error) {
	response := new(AuthLogoutResponse)
	c.Post("logout", nil, response)
	return response, nil
}

func (c *Client) AuthGoogleLogin(req *AuthGoogleLoginRequest) (*AuthLoginResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(AuthLoginResponse)
	err = c.Post("login", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) AuthFacebookLogin(req *AuthFacebookLoginRequest) (*AuthLoginResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(AuthLoginResponse)
	err = c.Post("login", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) AuthTwitterLogin(req *AuthTwitterLoginRequest) (*AuthLoginResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(AuthLoginResponse)
	err = c.Post("login", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) AuthMe() (*AuthMeResponse, error) {
	response := new(AuthMeResponse)
	if err := c.Get("me", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
