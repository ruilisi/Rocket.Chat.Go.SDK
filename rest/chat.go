package rest

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type ChatSyncThreadListResponse struct {
	Status
	Threads struct {
		Update []struct {
			ID  string    `json:"_id"`
			Rid string    `json:"rid"`
			Msg string    `json:"msg"`
			Ts  time.Time `json:"ts"`
			U   struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"u"`
			UpdatedAt time.Time     `json:"_updatedAt"`
			Mentions  []interface{} `json:"mentions"`
			Channels  []interface{} `json:"channels"`
			Replies   []string      `json:"replies"`
			Tcount    int           `json:"tcount"`
			Tlm       time.Time     `json:"tlm"`
		} `json:"update"`
		Remove []struct {
			ID  string    `json:"_id"`
			Rid string    `json:"rid"`
			Msg string    `json:"msg"`
			Ts  time.Time `json:"ts"`
			U   struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"u"`
			UpdatedAt  time.Time     `json:"_updatedAt"`
			Mentions   []interface{} `json:"mentions"`
			Channels   []interface{} `json:"channels"`
			Replies    []string      `json:"replies"`
			Tcount     int           `json:"tcount"`
			Tlm        time.Time     `json:"tlm"`
			DeletedAt  time.Time     `json:"_deletedAt"`
			Collection string        `json:"__collection__"`
		} `json:"remove"`
	} `json:"threads"`
}

type ChatSyncThreadResponse struct {
	Status
	Messages struct {
		Update []struct {
			ID   string    `json:"_id"`
			Rid  string    `json:"rid"`
			Tmid string    `json:"tmid"`
			Msg  string    `json:"msg"`
			Ts   time.Time `json:"ts"`
			U    struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"u"`
			UpdatedAt time.Time     `json:"_updatedAt"`
			Mentions  []interface{} `json:"mentions"`
			Channels  []interface{} `json:"channels"`
		} `json:"update"`
		Remove []interface{} `json:"remove"`
	} `json:"messages"`
}

type ChatReportResponse struct {
	Status
}

type ChatReactResponse struct {
	Status
}

type ChatGetThreadsListResponse struct {
	Status
	Threads []struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		UpdatedAt time.Time     `json:"_updatedAt"`
		Mentions  []interface{} `json:"mentions"`
		Channels  []interface{} `json:"channels"`
		Replies   []string      `json:"replies"`
		Tcount    int           `json:"tcount"`
		Tlm       time.Time     `json:"tlm"`
	} `json:"threads"`
}

type ChatGetThreadResponse struct {
	Status
	Messages []struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		UpdatedAt time.Time     `json:"_updatedAt"`
		Mentions  []interface{} `json:"mentions"`
		Channels  []interface{} `json:"channels"`
		Replies   []string      `json:"replies,omitempty"`
		Tcount    int           `json:"tcount,omitempty"`
		Tlm       time.Time     `json:"tlm,omitempty"`
		Tmid      string        `json:"tmid,omitempty"`
	} `json:"messages"`
}

type ChatGetSnippetedResponse struct {
	Status
	Messages []struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		UpdatedAt   time.Time     `json:"_updatedAt"`
		Mentions    []interface{} `json:"mentions"`
		Channels    []interface{} `json:"channels"`
		SnippetName string        `json:"snippetName"`
		Snippeted   bool          `json:"snippeted"`
		SnippetedBy struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"snippetedBy"`
	} `json:"messages"`
}

type ChatGetReadReceiptsResponse struct {
	Status
	Receipts []struct {
		ID        string    `json:"_id"`
		RoomID    string    `json:"roomId"`
		UserID    string    `json:"userId"`
		MessageID string    `json:"messageId"`
		Ts        time.Time `json:"ts"`
		User      struct {
			Username string `json:"username"`
			Name     string `json:"name"`
			ID       string `json:"_id"`
		} `json:"user"`
	} `json:"receipts"`
}

type ChatGetMentionedResponse struct {
	Status
	Messages []struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		UpdatedAt time.Time `json:"_updatedAt"`
		Mentions  []struct {
			ID       string `json:"_id"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"mentions"`
		Channels []interface{} `json:"channels"`
	} `json:"messages"`
}

type ChatGetDiscussionResponse struct {
	Status
	Messages []struct {
		ID  string    `json:"_id"`
		T   string    `json:"t"`
		Rid string    `json:"rid"`
		Ts  time.Time `json:"ts"`
		Msg string    `json:"msg"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Groupable   bool          `json:"groupable"`
		Drid        string        `json:"drid"`
		Attachments []interface{} `json:"attachments"`
		UpdatedAt   time.Time     `json:"_updatedAt"`
	} `json:"messages"`
}

type ChatGetDeletedResponse struct {
	Status
	Messages []struct {
		ID string `json:"_id"`
	} `json:"messages"`
}

type ChatSearchResponse struct {
	Status
	Messages []struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		Mentions  []interface{} `json:"mentions"`
		Channels  []interface{} `json:"channels"`
		UpdatedAt time.Time     `json:"_updatedAt"`
		Score     float64       `json:"score"`
	} `json:"messages"`
}

type ChatSendMessageResponse struct {
	Status
	Message struct {
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		Unread    bool          `json:"unread"`
		Mentions  []interface{} `json:"mentions"`
		Channels  []interface{} `json:"channels"`
		UpdatedAt time.Time     `json:"_updatedAt"`
		ID        string        `json:"_id"`
	} `json:"message"`
}

type ChatUnFollowResponse struct {
	Status
}

type ChatUnStarResponse struct {
	Status
}

type ChatUnPinResponse struct {
	Status
}

type ChatGetStarResponse struct {
	Status
	Messages []struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		UpdatedAt time.Time     `json:"_updatedAt"`
		Mentions  []interface{} `json:"mentions"`
		Channels  []interface{} `json:"channels"`
		Starred   []struct {
			ID string `json:"_id"`
		} `json:"starred"`
	} `json:"messages"`
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ChatGetPinResponse struct {
	Status
	Messages []struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		Mentions  []interface{} `json:"mentions"`
		Channels  []interface{} `json:"channels"`
		UpdatedAt time.Time     `json:"_updatedAt"`
		Pinned    bool          `json:"pinned"`
		PinnedAt  time.Time     `json:"pinnedAt"`
		PinnedBy  struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"pinnedBy"`
	} `json:"messages"`
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ChatStarResponse struct {
	Status
}

type ChatFollowResponse struct {
	Status
}

type ChatPinResponse struct {
	Status
	Message struct {
		T   string    `json:"t"`
		Rid string    `json:"rid"`
		Ts  time.Time `json:"ts"`
		Msg string    `json:"msg"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Groupable   bool `json:"groupable"`
		Attachments []struct {
			Text       string    `json:"text"`
			AuthorName string    `json:"author_name"`
			AuthorIcon string    `json:"author_icon"`
			Ts         time.Time `json:"ts"`
		} `json:"attachments"`
		UpdatedAt time.Time `json:"_updatedAt"`
		ID        string    `json:"_id"`
	} `json:"message"`
}

type ChatDeleteResponse struct {
	Status
	ID string `json:"_id"`
	Ts int64  `json:"ts"`
}
type ChatUpdateResponse struct {
	Status
	ChatUpdateMessage `json:"message"`
}

type ChatUpdateU struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
}
type ChatUpdateEditedBy struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
}
type ChatUpdateMessage struct {
	ID        string             `json:"_id"`
	Rid       string             `json:"rid"`
	Msg       string             `json:"msg"`
	Ts        time.Time          `json:"ts"`
	U         ChatUpdateU        `json:"u"`
	UpdatedAt time.Time          `json:"_updatedAt"`
	EditedAt  time.Time          `json:"editedAt"`
	EditedBy  ChatUpdateEditedBy `json:"editedBy"`
}

type ChatGetMessageResponse struct {
	Status
	models.CHATMessage `json:"message"`
}

type ChatIgnoreResponse struct {
	Status
}

func (c *Client) ChatDelete(req *models.ChatDeleteRequest) (*ChatDeleteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(ChatDeleteResponse)
	err = c.Post("chat.delete", bytes.NewBuffer(body), response)
	return response, err

}

func (c *Client) ChatGetMessage(params url.Values) (*models.CHATMessage, error) {

	response := new(ChatGetMessageResponse)
	if err := c.Get("chat.getMessage", params, response); err != nil {
		return nil, err
	}

	return &response.CHATMessage, nil
}

func (c *Client) ChatUpdate(req *models.ChatUpdateRequest) (*ChatUpdateResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(ChatUpdateResponse)
	err = c.Post("chat.update", bytes.NewBuffer(body), response)
	return response, nil

}

func (c *Client) ChatIgnore(params url.Values) (*ChatIgnoreResponse, error) {

	response := new(ChatIgnoreResponse)
	if err := c.Get("chat.ignoreUser", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatStar(req *models.ChatStarRequest) (*ChatStarResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(ChatStarResponse)
	err = c.Post("chat.starMessage", bytes.NewBuffer(body), response)
	return response, nil

}

func (c *Client) ChatPin(req *models.ChatPinRequest) (*ChatPinResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChatPinResponse)
	err = c.Post("chat.pinMessage", bytes.NewBuffer(body), response)
	return response, nil

}

func (c *Client) ChatFollow(req *models.ChatFollowRequest) (*ChatFollowResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChatFollowResponse)
	err = c.Post("chat.followMessage", bytes.NewBuffer(body), response)
	return response, nil

}

func (c *Client) ChatGetPin(params url.Values) (*ChatGetPinResponse, error) {

	response := new(ChatGetPinResponse)
	if err := c.Get("chat.getPinnedMessages", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatGetStar(params url.Values) (*ChatGetStarResponse, error) {

	response := new(ChatGetStarResponse)
	if err := c.Get("chat.getStarredMessages", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatUnPin(req *models.ChatUnPinRequest) (*ChatUnPinResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChatUnPinResponse)
	err = c.Post("chat.unPinMessage", bytes.NewBuffer(body), response)
	return response, nil

}

func (c *Client) ChatUnStar(req *models.ChatUnStarRequest) (*ChatUnStarResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChatUnStarResponse)
	err = c.Post("chat.unStarMessage", bytes.NewBuffer(body), response)
	return response, nil

}

func (c *Client) ChatUnFollow(req *models.ChatUnFollowRequest) (*ChatUnFollowResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChatUnFollowResponse)
	err = c.Post("chat.unfollowMessage", bytes.NewBuffer(body), response)
	return response, nil

}

func (c *Client) ChatSendMessage(req *models.ChatSendMessageRequest) (*ChatSendMessageResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChatSendMessageResponse)
	err = c.Post("chat.sendMessage", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChatSearch(params url.Values) (*ChatSearchResponse, error) {

	response := new(ChatSearchResponse)
	if err := c.Get("chat.search", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatGetDeleted(params url.Values) (*ChatGetDeletedResponse, error) {

	response := new(ChatGetDeletedResponse)
	if err := c.Get("chat.getDeletedMessages", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatGetDiscusssion(params url.Values) (*ChatGetDiscussionResponse, error) {

	response := new(ChatGetDiscussionResponse)
	if err := c.Get("chat.getDiscussions", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatGetMentioned(params url.Values) (*ChatGetMentionedResponse, error) {

	response := new(ChatGetMentionedResponse)
	if err := c.Get("chat.getMentionedMessages", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatGetReadReceipts(params url.Values) (*ChatGetReadReceiptsResponse, error) {

	response := new(ChatGetReadReceiptsResponse)
	if err := c.Get("chat.getMessageReadReceipts", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatGetSnippeted(params url.Values) (*ChatGetSnippetedResponse, error) {

	response := new(ChatGetSnippetedResponse)
	if err := c.Get("chat.getSnippetedMessages", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatGetThread(params url.Values) (*ChatGetThreadResponse, error) {

	response := new(ChatGetThreadResponse)
	if err := c.Get("chat.getThreadMessages", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatGetThreadsList(params url.Values) (*ChatGetThreadsListResponse, error) {

	response := new(ChatGetThreadsListResponse)
	if err := c.Get("chat.getThreadsList", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatReact(req *models.ChatReactRequest) (*ChatReactResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChatReactResponse)
	err = c.Post("chat.react", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChatReport(req *models.ChatReportRequest) (*ChatReportResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChatReportResponse)
	err = c.Post("chat.reportMessage", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChatSyncThread(params url.Values) (*ChatSyncThreadResponse, error) {

	response := new(ChatSyncThreadResponse)
	if err := c.Get("chat.syncThreadMessages", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChatSyncThreadList(params url.Values) (*ChatSyncThreadListResponse, error) {

	response := new(ChatSyncThreadListResponse)
	if err := c.Get("chat.syncThreadsList", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
