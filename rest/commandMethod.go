package rest

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type CommandGetResponse struct {
	Command models.Command
	Status
}

type CommandListResponse struct {
	Commands []models.Command
	Status
}

type CommandPreviewGetReqeuset struct {
	Command string `json:"command"`
	RoomId  string `json:"roomId"`
	Params  string `json:"params"`
}

type CommandPreviewPostRequest struct {
	Command     string `json:"command"`
	RoomId      string `json:"roomId"`
	TmId        string `json:"tmid"`
	Params      string `json:"params"`
	TriggerId   string `json:"triggerId"`
	PreviewItem struct {
		Id    string `json:"id"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"previewItem"`
}

type CommandRunRequest struct {
	Command   string `json:"command"`
	RoomID    string `json:"roomId"`
	Params    string `json:"params,omitempty"`
	Tmid      string `json:"tmid,omitempty"`
	TriggerID string `json:"triggerId,omitempty"`
}

type CommandRunResponse struct {
	Status
}

func (c *Client) CommandGet(params url.Values) (*CommandGetResponse, error) {
	response := new(CommandGetResponse)
	if err := c.Get("commands.get", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) CommandList(params url.Values) (*CommandListResponse, error) {
	response := new(CommandListResponse)
	if err := c.Get("commands.list", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) CommandPreviewGet(params url.Values) {
	c.Get("commands.preview", params, nil)
}
func (c *Client) CommandPreviewPost(req *CommandPreviewPostRequest) {
	body, _ := json.Marshal(req)
	c.Post("commands.preview", bytes.NewBuffer(body), nil)
}

func (c *Client) CommandRun(req *CommandRunRequest) (*CommandRunResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(CommandRunResponse)
	err = c.Post("commands.run", bytes.NewBuffer(body), response)
	return response, err
}
