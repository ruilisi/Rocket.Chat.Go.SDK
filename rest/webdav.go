package rest

import "net/url"

type WebdavGetAccountResponse struct {
	Status
	Accounts []struct {
		ID        string `json:"_id"`
		ServerURL string `json:"server_url"`
		Username  string `json:"username"`
		Name      string `json:"name"`
	} `json:"accounts"`
}

func (c *Client) WebdavGetAccount(params url.Values) (*WebdavGetAccountResponse, error) {
	response := new(WebdavGetAccountResponse)
	if err := c.Get("webdav.getMyAccounts", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
