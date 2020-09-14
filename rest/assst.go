package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type AssetUnSetResponse struct {
	Status
}

type AssetSetResponse struct {
	Status
}

func (c *Client) AssetSet(params map[string]io.Reader) (*AssetSetResponse, error) {
	response := new(AssetSetResponse)
	if err := c.PostFormData(http.MethodPost, "assets.setAsset", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) AssetUnSet(req *models.AssetUnSetRequest) (*AssetUnSetResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(AssetUnSetResponse)
	err = c.Post("assets.unsetAsset", bytes.NewBuffer(body), response)
	return response, nil
}
