package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type UploadResponse struct {
	Status
	Message struct {
		ID   string    `json:"_id"`
		Rid  string    `json:"rid"`
		Ts   time.Time `json:"ts"`
		Msg  string    `json:"msg"`
		File struct {
			ID   string `json:"_id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"file"`
		Attachments []struct {
			Ts                time.Time `json:"ts"`
			Title             string    `json:"title"`
			TitleLink         string    `json:"title_link"`
			TitleLinkDownload bool      `json:"title_link_download"`
			Type              string    `json:"type"`
			Description       string    `json:"description"`
		} `json:"attachments"`
		U struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		UpdatedAt time.Time `json:"_updatedAt"`
	} `json:"message"`
}

type RoomAdminRoomResponse struct {
	Status
	Rooms []struct {
		ID         string        `json:"_id"`
		T          string        `json:"t"`
		Name       string        `json:"name"`
		Usernames  []interface{} `json:"usernames"`
		Msgs       int           `json:"msgs"`
		UsersCount int           `json:"usersCount"`
		Default    bool          `json:"default"`
	} `json:"rooms"`
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type RoomCleanHistoryResponse struct {
	Status
	Message struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
	} `json:"message"`
}

type GetRoomResponse struct {
	Status
	Update []struct {
		ID          string        `json:"_id"`
		T           string        `json:"t"`
		Name        string        `json:"name"`
		Usernames   []interface{} `json:"usernames,omitempty"`
		UsersCount  int           `json:"usersCount"`
		Default     bool          `json:"default"`
		UpdatedAt   time.Time     `json:"_updatedAt"`
		LastMessage struct {
			ID          string        `json:"_id"`
			Alias       string        `json:"alias"`
			Msg         string        `json:"msg"`
			Attachments []interface{} `json:"attachments"`
			ParseUrls   bool          `json:"parseUrls"`
			Groupable   bool          `json:"groupable"`
			Ts          time.Time     `json:"ts"`
			U           struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"u"`
			Rid       string        `json:"rid"`
			UpdatedAt time.Time     `json:"_updatedAt"`
			Mentions  []interface{} `json:"mentions"`
			Channels  []interface{} `json:"channels"`
		} `json:"lastMessage"`
		Fname string `json:"fname,omitempty"`
		U     struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u,omitempty"`
		CustomFields struct {
		} `json:"customFields,omitempty"`
		Broadcast bool `json:"broadcast,omitempty"`
		Encrypted bool `json:"encrypted,omitempty"`
		Ro        bool `json:"ro,omitempty"`
		SysMes    bool `json:"sysMes,omitempty"`
	} `json:"update"`
	Remove []interface{} `json:"remove"`
}

type GetRoomInfoResponse struct {
	Status
	Room struct {
		ID          string        `json:"_id"`
		Ts          time.Time     `json:"ts"`
		T           string        `json:"t"`
		Name        string        `json:"name"`
		Usernames   []interface{} `json:"usernames"`
		Msgs        int           `json:"msgs"`
		UsersCount  int           `json:"usersCount"`
		Default     bool          `json:"default"`
		UpdatedAt   time.Time     `json:"_updatedAt"`
		LastMessage struct {
			ID          string        `json:"_id"`
			Alias       string        `json:"alias"`
			Msg         string        `json:"msg"`
			Attachments []interface{} `json:"attachments"`
			ParseUrls   bool          `json:"parseUrls"`
			Groupable   bool          `json:"groupable"`
			Ts          time.Time     `json:"ts"`
			U           struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"u"`
			Rid       string        `json:"rid"`
			UpdatedAt time.Time     `json:"_updatedAt"`
			Mentions  []interface{} `json:"mentions"`
			Channels  []interface{} `json:"channels"`
		} `json:"lastMessage"`
		Lm time.Time `json:"lm"`
	} `json:"room"`
}

type GetRoomDiscussionResponse struct {
	Status
	Discussions []struct {
		ID         string `json:"_id"`
		Name       string `json:"name"`
		Fname      string `json:"fname"`
		T          string `json:"t"`
		Msgs       int    `json:"msgs"`
		UsersCount int    `json:"usersCount"`
		U          struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Topic       string    `json:"topic"`
		Prid        string    `json:"prid"`
		Ts          time.Time `json:"ts"`
		Ro          bool      `json:"ro"`
		Default     bool      `json:"default"`
		SysMes      bool      `json:"sysMes"`
		UpdatedAt   time.Time `json:"_updatedAt"`
		LastMessage struct {
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
		} `json:"lastMessage"`
		Lm time.Time `json:"lm"`
	} `json:"discussions"`
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type CreateRoomResponse struct {
	Status
}

type LeaveRoomResponse struct {
	Status
}

type FavoriteResponse struct {
	Status
}

type SaveNotification struct {
	Status
}

func (c *Client) GetRoom(params url.Values) (*GetRoomResponse, error) {
	response := new(GetRoomResponse)
	if err := c.Get("rooms.get", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) RoomInfo(params url.Values) (*GetRoomInfoResponse, error) {
	response := new(GetRoomInfoResponse)
	if err := c.Get("rooms.info", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateDiscussion(discussionName string, room *models.Room) error {
	var body = fmt.Sprintf(`{ "prid": "%v","t_name":"%v"}`, room.ID, discussionName)
	return c.Post("rooms.createDiscussion", bytes.NewBufferString(body), new(CreateRoomResponse))
}

func (c *Client) GetDiscussion(params url.Values) (*GetRoomDiscussionResponse, error) {
	response := new(GetRoomDiscussionResponse)
	if err := c.Get("rooms.getDiscussions", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) LeaveRoom(room *models.Room) error {
	var body = fmt.Sprintf(`{ "roomId": "%v"}`, room.ID)
	return c.Post("rooms.leave", bytes.NewBufferString(body), new(LeaveRoomResponse))
}

func (c *Client) FavoriteRoom(room *models.Room) error {
	str := map[string]interface{}{
		"roomId":   room.ID,
		"favorite": room.Favorite,
	}
	data, _ := json.Marshal(str)
	return c.Post("rooms.favorite", bytes.NewReader(data), new(FavoriteResponse))
}

func (c *Client) SaveNotification(room *models.Room, notification *models.Notifications) error {
	var Notification map[string]interface{}
	noti, _ := json.Marshal(notification)
	json.Unmarshal(noti, &Notification)

	str := map[string]interface{}{
		"roomId":        room.ID,
		"notifications": Notification,
		// "notifications": structs.Map(notification),
	}
	data, _ := json.Marshal(str)
	return c.Post("rooms.saveNotification", bytes.NewReader(data), new(SaveNotification))
}

func (c *Client) RoomCleanHistory(req *models.RoomCleanHistoryRequest) (*RoomCleanHistoryResponse, error) {
	body, err := json.Marshal(req)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	response := new(RoomCleanHistoryResponse)
	err = c.Post("rooms.cleanHistory", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) RoomAdminRoom(params url.Values) (*RoomAdminRoomResponse, error) {

	response := new(RoomAdminRoomResponse)
	if err := c.Get("rooms.adminRooms", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Upload(params map[string]io.Reader, room *models.Room) (*UploadResponse, error) {
	response := new(UploadResponse)
	if err := c.PostFormData(http.MethodPost, "rooms.upload"+"/"+room.ID, params, response); err != nil {
		return nil, err
	}
	return response, nil
}
