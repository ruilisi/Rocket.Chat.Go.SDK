package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type base struct {
	Status
	Count  int
	Offset int
	Total  int
}

type CreateIMResponse struct {
	Status
	Room struct {
		Rid       string   `json:"rid"`
		T         string   `json:"t"`
		Usernames []string `json:"usernames"`
		ID        string   `json:"_id"`
	} `json:"room"`
}

type CountersIMResponse struct {
	Status
	models.CountersIMMessage
}

type CloseIMResponse struct {
	Status
}

type HistoryIMResponse struct {
	Status
	Messages []models.Message `json:"messages"`
}

type FileIMResponse struct {
	base
	Files []models.FileIM
}

type MemberIMResponse struct {
	base
	Members []models.MemberIM
}

type MessageIMResponse struct {
	base
	Messages []models.Message
}

type OtherMessageIMResponse struct {
	Status
	Messages []models.Message `json:"messages"`
}

type ListIMResponse struct {
	base
	Ims []models.ListIM
}

type EveryoneListIMResponse struct {
	base
	Ims []models.EveryoneListIM
}

type OpenIM struct {
	RoomId string `json:"roomId"`
}

type OpenIMResponse struct {
	Status
}

type SetTopic struct {
	RoomId string `json:"roomId"`
	Topic  string `json:"topic"`
}

type SetTopicIMResponse struct {
	Topic string
	Status
}

func (c *Client) CreateIM(req *models.CreateIMRequest) (*CreateIMResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(CreateIMResponse)
	err = c.Post("im.create", bytes.NewBuffer(body), response)
	return response, err

}

func (c *Client) CountersIM(params url.Values) (*models.CountersIMMessage, error) {

	response := new(CountersIMResponse)
	if err := c.Get("im.counters", params, response); err != nil {
		return nil, err
	}

	return &response.CountersIMMessage, nil
}

func (c *Client) CloseIM(req *models.CloseIMRequest) (*CloseIMResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(CloseIMResponse)
	err = c.Post("im.counters", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) HistoryIM(params url.Values) (*HistoryIMResponse, error) {
	response := new(HistoryIMResponse)
	if err := c.Get("im.history", params, response); err != nil {
		return nil, err
	}
	return response, nil

}
func (c *Client) FilesIM(params url.Values) (*FileIMResponse, error) {
	response := new(FileIMResponse)
	if err := c.Get("im.files", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) MembersIM(params url.Values) (*MemberIMResponse, error) {
	response := new(MemberIMResponse)
	if err := c.Get("im.members", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) MessagesIM(params url.Values) (*MessageIMResponse, error) {
	response := new(MessageIMResponse)
	if err := c.Get("im.messages", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) OtherMessagesIM(params url.Values) (*OtherMessageIMResponse, error) {
	response := new(OtherMessageIMResponse)
	if err := c.Get("im.messages.others", params, response); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response, nil
}

func (c *Client) ListIM() (*ListIMResponse, error) {
	response := new(ListIMResponse)
	if err := c.Get("im.list", nil, response); err != nil {
		return nil, err
	}
	return response, nil

}

func (c *Client) EveryoneListIM() (*EveryoneListIMResponse, error) {
	response := new(EveryoneListIMResponse)
	if err := c.Get("im.list.everyone", nil, response); err != nil {
		return nil, err
	}
	return response, nil

}

func (c *Client) OpenIM(req *OpenIM) (*OpenIMResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(OpenIMResponse)
	err = c.Post("im.open", bytes.NewBuffer(body), response)
	return response, err

}

func (c *Client) SetTopicIM(req *SetTopic) (*SetTopicIMResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(SetTopicIMResponse)
	err = c.Post("im.setTopic", bytes.NewBuffer(body), response)
	return response, err

}
