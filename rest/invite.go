package rest

import (
	"bytes"
	"encoding/json"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type CreateInviteResponse struct {
	Status
	models.Invite
}

type ListInvitesResponse []models.Invite

type RemoveInviteRequest struct {
	Id string `json:"_id"`
}

type RemoveInviteResponse bool

type UserInviteToken struct {
	Token string `json:"token"`
}

type UserInviteTokenResponse struct {
	Room struct {
		Rid   string `json:"rid"`
		Fname string `json:"fname"`
		Name  string `json:"name"`
		T     string `json:"t"`
	} `json:"room"`
	Status
}

type ValidateInviteTokenResponse struct {
	Valid bool `json:"valid"`
	Status
}

func (c *Client) CreateInvite(req *models.CreateInviteRequest) (*CreateInviteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(CreateInviteResponse)
	err = c.Post("findOrCreateInvite", bytes.NewBuffer(body), response)
	return response, err

}

func (c *Client) ListInvites() (*ListInvitesResponse, error) {
	response := new(ListInvitesResponse)
	if err := c.nGet("listInvites", nil, response); err != nil {
		return nil, err
	}

	return response, nil

}

func (c *Client) RemoveInvite(req *RemoveInviteRequest) (RemoveInviteResponse, error) {
	api := "removeInvite/" + req.Id
	response := new(RemoveInviteResponse)
	if err := c.nDelete(api, nil, response); err != nil {
		return false, err
	}

	return *response, nil
}

func (c *Client) UserInviteToken(req *UserInviteToken) (*UserInviteTokenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(UserInviteTokenResponse)
	err = c.Post("useInviteToken", bytes.NewBuffer(body), response)
	return response, err

}

func (c *Client) ValidateInviteToken(req *UserInviteToken) (*ValidateInviteTokenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(ValidateInviteTokenResponse)
	err = c.Post("validateInviteToken", bytes.NewBuffer(body), response)
	return response, err

}
