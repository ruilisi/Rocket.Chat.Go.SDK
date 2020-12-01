package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type ChannelAnonymousreadResponse struct {
	Status
	Messages []struct {
		ID          string        `json:"_id"`
		Alias       string        `json:"alias"`
		Msg         string        `json:"msg"`
		Attachments []interface{} `json:"attachments"`
		ParseUrls   bool          `json:"parseUrls,omitempty"`
		Bot         struct {
			I string `json:"i"`
		} `json:"bot,omitempty"`
		Groupable bool      `json:"groupable,omitempty"`
		Ts        time.Time `json:"ts"`
		U         struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		Rid       string    `json:"rid"`
		UpdatedAt time.Time `json:"_updatedAt"`
		Reactions struct {
			Grin struct {
				Usernames []string `json:"usernames"`
			} `json:":grin:"`
		} `json:"reactions,omitempty"`
		Mentions []interface{} `json:"mentions"`
		Channels []interface{} `json:"channels"`
		Starred  struct {
			ID string `json:"_id"`
		} `json:"starred,omitempty"`
		Emoji    string `json:"emoji,omitempty"`
		Avatar   string `json:"avatar,omitempty"`
		EditedBy struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"editedBy,omitempty"`
		EditedAt time.Time     `json:"editedAt,omitempty"`
		Urls     []interface{} `json:"urls,omitempty"`
	} `json:"messages"`
}

type ChannelCustomFieldsResponse struct {
	Status
	Channel struct {
		ID           string    `json:"_id"`
		Ts           time.Time `json:"ts"`
		T            string    `json:"t"`
		Name         string    `json:"name"`
		Msgs         int       `json:"msgs"`
		Default      bool      `json:"default"`
		UpdatedAt    time.Time `json:"_updatedAt"`
		Lm           time.Time `json:"lm"`
		CustomFields struct {
			Organization string `json:"organization"`
		} `json:"customFields"`
	} `json:"channel"`
	DeveloperWarning string `json:"developerWarning"`
}

type ChannelMentionResponse struct {
	Status
	Mentions []struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		Mentions []struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"mentions"`
		Channels  []interface{} `json:"channels"`
		UpdatedAt time.Time     `json:"_updatedAt"`
	} `json:"mentions"`
}

type ChannelOpenResponse struct {
	Status
}

type ChannelHistoryResponse struct {
	Status
	Messages []struct {
		ID  string    `json:"_id"`
		Rid string    `json:"rid"`
		Msg string    `json:"msg"`
		Ts  time.Time `json:"ts"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		UpdatedAt time.Time `json:"_updatedAt"`
		T         string    `json:"t,omitempty"`
		Groupable bool      `json:"groupable,omitempty"`
	} `json:"messages"`
}

type ChannelDefaultResponse struct {
	Status
	Channel struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
		T    string `json:"t"`
		Msgs int    `json:"msgs"`
		U    struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Ts               time.Time     `json:"ts"`
		Ro               bool          `json:"ro"`
		SysMes           bool          `json:"sysMes"`
		UpdatedAt        time.Time     `json:"_updatedAt"`
		Usernames        []string      `json:"usernames"`
		JoinCodeRequired bool          `json:"joinCodeRequired"`
		Muted            []interface{} `json:"muted"`
		Default          bool          `json:"default"`
	} `json:"channel"`
}

type ChannelRolesResponse struct {
	Status
	Roles []struct {
		Rid string `json:"rid"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		Roles []string `json:"roles"`
		ID    string   `json:"_id"`
	} `json:"roles"`
}

type ChannelKickResponse struct {
	Status
	Channel struct {
		ID        string   `json:"_id"`
		Name      string   `json:"name"`
		T         string   `json:"t"`
		Usernames []string `json:"usernames"`
		Msgs      int      `json:"msgs"`
		U         struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Ts        time.Time `json:"ts"`
		Ro        bool      `json:"ro"`
		SysMes    bool      `json:"sysMes"`
		UpdatedAt time.Time `json:"_updatedAt"`
	} `json:"channel"`
}

type ChannelJoinResponse struct {
	Status
	Channel struct {
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
		CustomFields struct {
		} `json:"customFields"`
		Broadcast        bool      `json:"broadcast"`
		Encrypted        bool      `json:"encrypted"`
		Ts               time.Time `json:"ts"`
		Ro               bool      `json:"ro"`
		SysMes           bool      `json:"sysMes"`
		Default          bool      `json:"default"`
		UpdatedAt        time.Time `json:"_updatedAt"`
		JoinCodeRequired bool      `json:"joinCodeRequired"`
	} `json:"channel"`
}

type ChannelDeleteResponse struct {
	Status
}

type ChannelCloseResponse struct {
	Status
}

type ChannelUnArchiveResponse struct {
	Status
}

type ChannelArchiveResponse struct {
	Status
}

type ChannelRenameResponse struct {
	Status
	Channel struct {
		ID        string   `json:"_id"`
		Name      string   `json:"name"`
		T         string   `json:"t"`
		Usernames []string `json:"usernames"`
		Msgs      int      `json:"msgs"`
		U         struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Ts        time.Time `json:"ts"`
		Ro        bool      `json:"ro"`
		SysMes    bool      `json:"sysMes"`
		UpdatedAt time.Time `json:"_updatedAt"`
	} `json:"channel"`
}

type ChannelAnnouncementResponse struct {
	Status
	Announcement string `json:"announcement"`
}

type ChannelTypeResponse struct {
	Status
	Channel struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
		T    string `json:"t"`
		Msgs int    `json:"msgs"`
		U    struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Ts               time.Time     `json:"ts"`
		Ro               bool          `json:"ro"`
		SysMes           bool          `json:"sysMes"`
		UpdatedAt        time.Time     `json:"_updatedAt"`
		Usernames        []string      `json:"usernames"`
		JoinCodeRequired bool          `json:"joinCodeRequired"`
		Muted            []interface{} `json:"muted"`
	} `json:"channel"`
}

type ChannelReadOnlyResponse struct {
	Status
	Channel struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
		T    string `json:"t"`
		Msgs int    `json:"msgs"`
		U    struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Ts               time.Time     `json:"ts"`
		Ro               bool          `json:"ro"`
		SysMes           bool          `json:"sysMes"`
		UpdatedAt        time.Time     `json:"_updatedAt"`
		Usernames        []string      `json:"usernames"`
		JoinCodeRequired bool          `json:"joinCodeRequired"`
		Muted            []interface{} `json:"muted"`
	} `json:"channel"`
}

type ChannelTopicResponse struct {
	Status
	Topic string `json:"topic"`
}

type ChannelJoinCodeResponse struct {
	Status
	Channel struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
		T    string `json:"t"`
		Msgs int    `json:"msgs"`
		U    struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Ts               time.Time `json:"ts"`
		Ro               bool      `json:"ro"`
		SysMes           bool      `json:"sysMes"`
		UpdatedAt        time.Time `json:"_updatedAt"`
		Usernames        []string  `json:"usernames"`
		JoinCodeRequired bool      `json:"joinCodeRequired"`
	} `json:"channel"`
}

type ChannelDescriptionResponse struct {
	Status
	Description string `json:"description"`
}

type ChannelModeratorsResponse struct {
	Status
	Moderators []struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
		Name     string `json:"name"`
	} `json:"moderators"`
}
type ChannelMessagesResponse struct {
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
	} `json:"messages"`
	Status
}

type ChannelMembersResponse struct {
	Status
	Members []struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Status   string `json:"status"`
	} `json:"members"`
}

type ChannelFilesResponse struct {
	Status
	Files []struct {
		ID          string `json:"_id"`
		Name        string `json:"name"`
		Size        int    `json:"size"`
		Type        string `json:"type"`
		Rid         string `json:"rid"`
		Description string `json:"description"`
		Store       string `json:"store"`
		Complete    bool   `json:"complete"`
		Uploading   bool   `json:"uploading"`
		Extension   string `json:"extension"`
		Progress    int    `json:"progress"`
		User        struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"user"`
		UpdatedAt  time.Time `json:"_updatedAt"`
		InstanceID string    `json:"instanceId"`
		Etag       string    `json:"etag"`
		Path       string    `json:"path"`
		Token      string    `json:"token"`
		UploadedAt time.Time `json:"uploadedAt"`
		URL        string    `json:"url"`
	} `json:"files"`
}

type ChannelCountersResponse struct {
	Status
	Joined       bool      `json:"joined"`
	Members      int       `json:"members"`
	Unreads      int       `json:"unreads"`
	UnreadsFrom  time.Time `json:"unreadsFrom"`
	Msgs         int       `json:"msgs"`
	Latest       time.Time `json:"latest"`
	UserMentions int       `json:"userMentions"`
}

type ChannelInviteResponse struct {
	Status
	Channel struct {
		ID        string    `json:"_id"`
		Ts        time.Time `json:"ts"`
		T         string    `json:"t"`
		Name      string    `json:"name"`
		Usernames []string  `json:"usernames"`
		Msgs      int       `json:"msgs"`
		UpdatedAt time.Time `json:"_updatedAt"`
		Lm        time.Time `json:"lm"`
	} `json:"channel"`
}

type ChannelRemoveModeratorResponse struct {
	Status
}

type ChannelAddModeratorResponse struct {
	Status
}

type ChannelRemoveOwnerResponse struct {
	Status
}

type ChannelAddOwnerResponse struct {
	Status
}

type ChannelRemoveLeaderResponse struct {
	Status
}

type ChannelAddLeaderResponse struct {
	Status
}

type ChannelAddAllResponse struct {
	Status
	Channel struct {
		ID        string   `json:"_id"`
		Name      string   `json:"name"`
		T         string   `json:"t"`
		Usernames []string `json:"usernames"`
		Msgs      int      `json:"msgs"`
		U         struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Ts time.Time `json:"ts"`
	} `json:"channel"`
}

type ChannelCreateResponse struct {
	Status
	Channel struct {
		ID        string   `json:"_id"`
		Name      string   `json:"name"`
		T         string   `json:"t"`
		Usernames []string `json:"usernames"`
		Msgs      int      `json:"msgs"`
		U         struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Ts time.Time `json:"ts"`
	} `json:"channel"`
}

type ChannelsResponse struct {
	Status
	models.Pagination
	Channels []models.Channel `json:"channels"`
}

type ChannelResponse struct {
	Status
	Channel models.Channel `json:"channel"`
}

type OnlineChannelResponse struct {
	Status
	Online []models.OnlineChannel `json:"online"`
}

// GetPublicChannels returns all channels that can be seen by the logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list
func (c *Client) ChannelsList(params url.Values) (*ChannelsResponse, error) {
	response := new(ChannelsResponse)
	if err := c.Get("channels.list", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetOnlineChannels(params url.Values) (*OnlineChannelResponse, error) {
	response := new(OnlineChannelResponse)
	if err := c.Get("channels.online", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetJoinedChannels returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list-joined
func (c *Client) GetJoinedChannels(params url.Values) (*ChannelsResponse, error) {
	response := new(ChannelsResponse)
	if err := c.Get("channels.list.joined", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// LeaveChannel leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *Client) LeaveChannel(channel *models.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.ID)
	return c.Post("channels.leave", bytes.NewBufferString(body), new(ChannelResponse))
}

// GetChannelInfo get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (c *Client) GetChannelInfo(channel *models.Channel) (*models.Channel, error) {
	response := new(ChannelResponse)
	switch {
	case channel.Name != "" && channel.ID == "":
		if err := c.Get("channels.info", url.Values{"roomName": []string{channel.Name}}, response); err != nil {
			return nil, err
		}
	default:
		if err := c.Get("channels.info", url.Values{"roomId": []string{channel.ID}}, response); err != nil {
			return nil, err
		}
	}

	return &response.Channel, nil
}

func (c *Client) ChannelCreate(req *models.ChannelCreateRequest) (*ChannelCreateResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelCreateResponse)
	err = c.Post("channels.create", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelAddAll(req *models.ChannelAddAllRequest) (*ChannelAddAllResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelAddAllResponse)
	err = c.Post("channels.addAll", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelAddLeader(req *models.ChannelAddLeaderRequest) (*ChannelAddLeaderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelAddLeaderResponse)
	err = c.Post("channels.addLeader", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelRemoveLeader(req *models.ChannelRemoveLeaderRequest) (*ChannelRemoveLeaderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelRemoveLeaderResponse)
	err = c.Post("channels.removeLeader", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelAddOwner(req *models.ChannelAddOwnerRequest) (*ChannelAddOwnerResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelAddOwnerResponse)
	err = c.Post("channels.addOwner", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelRemoveOwner(req *models.ChannelRemoveOwnerRequest) (*ChannelRemoveOwnerResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelRemoveOwnerResponse)
	err = c.Post("channels.removeOwner", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelAddModerator(req *models.ChannelAddModeratorRequest) (*ChannelAddModeratorResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelAddModeratorResponse)
	err = c.Post("channels.addModerator", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelRemoveModerator(req *models.ChannelRemoveModeratorRequest) (*ChannelRemoveModeratorResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelRemoveModeratorResponse)
	err = c.Post("channels.removeModerator", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelInvite(req *models.ChannelInviteRequest) (*ChannelInviteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelInviteResponse)
	err = c.Post("channels.invite", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelCounters(params url.Values) (*ChannelCountersResponse, error) {

	response := new(ChannelCountersResponse)
	if err := c.Get("channels.counters", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChannelFiles(params url.Values) (*ChannelFilesResponse, error) {

	response := new(ChannelFilesResponse)
	if err := c.Get("channels.files", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChannelMembers(params url.Values) (*ChannelMembersResponse, error) {

	response := new(ChannelMembersResponse)
	if err := c.Get("channels.members", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChannelMessages(params url.Values) (*ChannelMessagesResponse, error) {

	response := new(ChannelMessagesResponse)
	if err := c.Get("channels.messages", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChannelModerators(params url.Values) (*ChannelModeratorsResponse, error) {

	response := new(ChannelModeratorsResponse)
	if err := c.Get("channels.moderators", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChannelDescription(req *models.ChannelDescriptionRequest) (*ChannelDescriptionResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelDescriptionResponse)
	err = c.Post("channels.setDescription", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelJoinCode(req *models.ChannelJoinCodeRequest) (*ChannelJoinCodeResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelJoinCodeResponse)
	err = c.Post("channels.setJoinCode", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelTopic(req *models.ChannelTopicRequest) (*ChannelTopicResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelTopicResponse)
	err = c.Post("channels.setTopic", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelReadOnly(req *models.ChannelReadOnlyRequest) (*ChannelReadOnlyResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelReadOnlyResponse)
	err = c.Post("channels.setReadOnly", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelType(req *models.ChannelTypeRequest) (*ChannelTypeResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelTypeResponse)
	err = c.Post("channels.setType", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelAnnouncement(req *models.ChannelAnnouncementRequest) (*ChannelAnnouncementResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelAnnouncementResponse)
	err = c.Post("channels.setAnnouncement", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelRename(req *models.ChannelRenameRequest) (*ChannelRenameResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelRenameResponse)
	err = c.Post("channels.rename", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelArchive(req *models.ChannelArchiveRequest) (*ChannelArchiveResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelArchiveResponse)
	err = c.Post("channels.archive", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelUnArchive(req *models.ChannelUnArchiveRequest) (*ChannelUnArchiveResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelUnArchiveResponse)
	err = c.Post("channels.unarchive", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelClose(req *models.ChannelCloseRequest) (*ChannelCloseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelCloseResponse)
	err = c.Post("channels.close", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelDelete(req *models.ChannelDeleteRequest) (*ChannelDeleteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelDeleteResponse)
	err = c.Post("channels.delete", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelJoin(req *models.ChannelJoinRequest) (*ChannelJoinResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelJoinResponse)
	err = c.Post("channels.join", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelKick(req *models.ChannelKickRequest) (*ChannelKickResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelKickResponse)
	err = c.Post("channels.kick", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelRoles(params url.Values) (*ChannelRolesResponse, error) {

	response := new(ChannelRolesResponse)
	if err := c.Get("channels.roles", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChannelDefault(req *models.ChannelDefaultRequest) (*ChannelDefaultResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelDefaultResponse)
	err = c.Post("channels.setDefault", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelOpen(req *models.ChannelOpenRequest) (*ChannelOpenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelOpenResponse)
	err = c.Post("channels.open", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelMention(params url.Values) (*ChannelMentionResponse, error) {

	response := new(ChannelMentionResponse)
	if err := c.Get("channels.getAllUserMentionsByChannel", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChannelCustomFields(req *models.ChannelCustomFieldsRequest) (*ChannelCustomFieldsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(ChannelCustomFieldsResponse)
	err = c.Post("channels.setCustomFields", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) ChannelAnonymousread(params url.Values) (*ChannelAnonymousreadResponse, error) {

	response := new(ChannelAnonymousreadResponse)
	if err := c.Get("channels.anonymousread", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ChannelHistory(params url.Values) (*ChannelHistoryResponse, error) {
	response := new(ChannelHistoryResponse)
	if err := c.Get("channels.history", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
