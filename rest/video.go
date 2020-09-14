package rest

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type VedieUpdateTimeoutResponse struct {
	Status
	JitsiTimeout int64 `json:"jitsiTimeout"`
}

func (c *Client) VedieUpdateTimeout(req *models.VedieUpdateTimeoutRequest) (*VedieUpdateTimeoutResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	response := new(VedieUpdateTimeoutResponse)
	err = c.Post("video-conference/jitsi.update-timeout", bytes.NewBuffer(body), response)
	return response, nil
}
