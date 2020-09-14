package rest

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type RolesListResponse struct {
	Status
	Roles []models.Roles `json:"roles"`
}

type RolesSyncResponse struct {
	Status
	Roles struct {
		Update []struct {
			ID           string    `json:"_id"`
			UpdatedAt    time.Time `json:"_updatedAt"`
			Description  string    `json:"description"`
			Mandatory2Fa bool      `json:"mandatory2fa"`
			Protected    bool      `json:"protected"`
			Scope        string    `json:"scope"`
		} `json:"update"`
		Remove []interface{} `json:"remove"`
	} `json:"roles"`
}

type AddUserToRoleResponse struct {
	Status
	Role struct {
		ID          string    `json:"_id"`
		Name        string    `json:"name"`
		Scope       string    `json:"scope"`
		Description string    `json:"description"`
		UpdatedAt   time.Time `json:"_updatedAt"`
	} `json:"role"`
}

type GetUsersInRoleResponse struct {
	Status
	Users []struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
		Type     string `json:"type"`
		Status   string `json:"status"`
		Active   bool   `json:"active"`
		Name     string `json:"name"`
	} `json:"users"`
}

type CreateRoleResponse struct {
	Status
	Role struct {
		ID          string    `json:"_id"`
		Name        string    `json:"name"`
		Scope       string    `json:"scope"`
		Description string    `json:"description"`
		UpdatedAt   time.Time `json:"_updatedAt"`
	} `json:"role"`
}

func (c *Client) RolesList() (*RolesListResponse, error) {
	response := new(RolesListResponse)
	if err := c.Get("roles.list", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) RolesSync(params url.Values) (*RolesSyncResponse, error) {
	response := new(RolesSyncResponse)
	if err := c.Get("roles.sync", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetUsersInRole(params url.Values) (*GetUsersInRoleResponse, error) {
	response := new(GetUsersInRoleResponse)
	if err := c.Get("roles.getUsersInRole", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateRole(req *models.Role) (*CreateRoleResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(CreateRoleResponse)
	err = c.Post("roles.create", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) AddUserToRole(req *models.AddUserToRole) (*AddUserToRoleResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(AddUserToRoleResponse)
	err = c.Post("roles.addUserToRole", bytes.NewBuffer(body), response)
	return response, err
}
