package rest

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type GroupCustomFieldsResponse struct {
	Status
	Group struct {
		ID    string `json:"_id"`
		Name  string `json:"name"`
		Fname string `json:"fname"`
		T     string `json:"t"`
		Msgs  int    `json:"msgs"`
		U     struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		CustomFields struct {
			Company string `json:"company"`
		} `json:"customFields"`
		Ts        time.Time `json:"ts"`
		Ro        bool      `json:"ro"`
		SysMes    bool      `json:"sysMes"`
		UpdatedAt time.Time `json:"_updatedAt"`
		Username  string    `json:"username"`
	} `json:"group"`
	DeveloperWarning string `json:"developerWarning"`
}

type GroupDeleteResponse struct {
	Status
}

type GroupArchiveResponse struct {
	Status
}

type GroupUnArchiveResponse struct {
	Status
}

type GroupAddAllResponse struct {
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

type GroupLeaveResponse struct {
	Status
}

type GroupHistoryResponse struct {
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

type GroupTypeResponse struct {
	Status
	Group struct {
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
	} `json:"group"`
}

type GroupInfoResponse struct {
	Status
	Group struct {
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
		Broadcast bool      `json:"broadcast"`
		Encrypted bool      `json:"encrypted"`
		Ts        time.Time `json:"ts"`
		Ro        bool      `json:"ro"`
		Default   bool      `json:"default"`
		SysMes    bool      `json:"sysMes"`
		UpdatedAt time.Time `json:"_updatedAt"`
	} `json:"group"`
}

type GroupRolesResponse struct {
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

type GroupOnlineResponse struct {
	Status
	Online []struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
	} `json:"online"`
}

type GroupTopicResponse struct {
	Status
	Topic string `json:"topic"`
}

type GroupReadOnlyResponse struct {
	Status
	Group struct {
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
	} `json:"group"`
}

type GroupDescriptionResponse struct {
	Status
	Description string `json:"description"`
}

type GroupAnnouncementResponse struct {
	Status
	Announcement string `json:"announcement"`
}

type GroupFilesResponse struct {
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

type GroupCountersResponse struct {
	Status
	Joined       bool      `json:"joined"`
	Members      int       `json:"members"`
	Unreads      int       `json:"unreads"`
	UnreadsFrom  time.Time `json:"unreadsFrom"`
	Msgs         int       `json:"msgs"`
	Latest       time.Time `json:"latest"`
	UserMentions int       `json:"userMentions"`
}

type GroupListAllResponse struct {
	Status
	Groups []struct {
		ID    string `json:"_id"`
		Name  string `json:"name"`
		Fname string `json:"fname"`
		T     string `json:"t"`
		Msgs  int    `json:"msgs"`
		U     struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		CustomFields struct {
			CompanyID string `json:"companyId"`
		} `json:"customFields"`
		Ts        time.Time `json:"ts"`
		Ro        bool      `json:"ro"`
		SysMes    bool      `json:"sysMes"`
		UpdatedAt time.Time `json:"_updatedAt"`
	} `json:"groups"`
}

type GroupModeratorsResponse struct {
	Status
	Moderators []struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
		Name     string `json:"name"`
	} `json:"moderators"`
}

type GroupMembersResponse struct {
	Status
	Members []struct {
		ID        string `json:"_id"`
		Status    string `json:"status"`
		Name      string `json:"name"`
		UtcOffset int    `json:"utcOffset"`
		Username  string `json:"username"`
	} `json:"members"`
}

type GroupMessagesResponse struct {
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
		} `json:"u,omitempty"`
		UpdatedAt time.Time `json:"_updatedAt"`
		Reactions struct {
			Frowning2 struct {
				Usernames []string `json:"usernames"`
			} `json:":frowning2:"`
		} `json:"reactions,omitempty"`
		Mentions []interface{} `json:"mentions,omitempty"`
		Channels []interface{} `json:"channels,omitempty"`
		Starred  struct {
			ID string `json:"_id"`
		} `json:"starred,omitempty"`
		T         string `json:"t,omitempty"`
		Groupable bool   `json:"groupable,omitempty"`
	} `json:"messages"`
}

type GroupKickResponse struct {
	Status
}

type GroupInviteResponse struct {
	Status
	Group struct {
		ID        string    `json:"_id"`
		Ts        time.Time `json:"ts"`
		T         string    `json:"t"`
		Name      string    `json:"name"`
		Usernames []string  `json:"usernames"`
		U         struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"u"`
		Msgs      int       `json:"msgs"`
		UpdatedAt time.Time `json:"_updatedAt"`
		Lm        time.Time `json:"lm"`
	} `json:"group"`
}

type GroupOpenResponse struct {
	Status
}

type GroupCloseResponse struct {
	Status
}

type GroupRenameResponse struct {
	Status
	Group struct {
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
	} `json:"group"`
}

type GroupRemoveOwnerResponse struct {
	Status
}

type GroupAddOwnerResponse struct {
	Status
}

type GroupRemoveModeratorResponse struct {
	Status
}

type GroupAddModeratorResponse struct {
	Status
}
type GroupAddLeaderResponse struct {
	Status
}

type GroupRemoveLeaderResponse struct {
	Status
}
type GroupListResponse struct {
	Status
	models.GroupList
}
type CreateGroupResponse struct {
	Status
	Group struct {
		ID        string            `json:"_id"`
		Name      string            `json:"name"`
		T         string            `json:"t"`
		UserNames map[string]string `json:"usernames"`
		Msgs      int               `json:"msgs"`
		U         map[string]string `json:"u"`
		Ts        time.Time         `json:"ts"`
		Ro        bool              `json:"ro"`
		SysMes    bool              `json:"sysMes"`
		UpdateAt  time.Time         `json:"_updatedAt"`
	} `json:"group"`
}

func (c *Client) CreateGroup(req *models.CreateGroupRequest) (*CreateGroupResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(CreateGroupResponse)
	err = c.Post("groups.create", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) GroupAddLeader(req *models.GroupAddLeaderRequest) (*GroupAddLeaderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(GroupAddLeaderResponse)
	err = c.Post("groups.addLeader", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) GroupRemoveLeader(req *models.GroupRemoveLeaderRequest) (*GroupRemoveLeaderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(GroupRemoveLeaderResponse)
	err = c.Post("groups.removeLeader", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) GroupList(params url.Values) (*GroupListResponse, error) {
	response := new(GroupListResponse)
	if err := c.Get("groups.list", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupAddModerator(req *models.GroupAddModeratorRequest) (*GroupAddModeratorResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupAddModeratorResponse)
	err = c.Post("groups.addModerator", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupRemoveModerator(req *models.GroupRemoveModeratorRequest) (*GroupRemoveModeratorResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupRemoveModeratorResponse)
	err = c.Post("groups.removeModerator", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupAddOwner(req *models.GroupAddOwnerRequest) (*GroupAddOwnerResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupAddOwnerResponse)
	err = c.Post("groups.addOwner", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupRemoveOwner(req *models.GroupRemoveOwnerRequest) (*GroupRemoveOwnerResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupRemoveOwnerResponse)
	err = c.Post("groups.removeOwner", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupRename(req *models.GroupRenameRequest) (*GroupRenameResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupRenameResponse)
	err = c.Post("groups.rename", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupClose(req *models.GroupCloseRequest) (*GroupCloseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupCloseResponse)
	err = c.Post("groups.close", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupOpen(req *models.GroupOpenRequest) (*GroupOpenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupOpenResponse)
	err = c.Post("groups.open", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupInvite(req *models.GroupInviteRequest) (*GroupInviteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupInviteResponse)
	err = c.Post("groups.invite", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupKick(req *models.GroupKickRequest) (*GroupKickResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupKickResponse)
	err = c.Post("groups.kick", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupMessages(params url.Values) (*GroupMessagesResponse, error) {

	response := new(GroupMessagesResponse)
	if err := c.Get("groups.messages", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupMembers(params url.Values) (*GroupMembersResponse, error) {

	response := new(GroupMembersResponse)
	if err := c.Get("groups.members", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupModerators(params url.Values) (*GroupModeratorsResponse, error) {

	response := new(GroupModeratorsResponse)
	if err := c.Get("groups.moderators", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupListAll(params url.Values) (*GroupListAllResponse, error) {

	response := new(GroupListAllResponse)
	if err := c.Get("groups.listAll", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupCounters(params url.Values) (*GroupCountersResponse, error) {

	response := new(GroupCountersResponse)
	if err := c.Get("groups.counters", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupFiles(params url.Values) (*GroupFilesResponse, error) {

	response := new(GroupFilesResponse)
	if err := c.Get("groups.files", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupAnnouncement(req *models.GroupAnnouncementRequest) (*GroupAnnouncementResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupAnnouncementResponse)
	err = c.Post("groups.setAnnouncement", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupDescription(req *models.GroupDescriptionRequest) (*GroupDescriptionResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupDescriptionResponse)
	err = c.Post("groups.setDescription", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupReadOnly(req *models.GroupReadOnlyRequest) (*GroupReadOnlyResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupReadOnlyResponse)
	err = c.Post("groups.setReadOnly", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupTopic(req *models.GroupTopicRequest) (*GroupTopicResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupTopicResponse)
	err = c.Post("groups.setTopic", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupOnline(params url.Values) (*GroupOnlineResponse, error) {

	response := new(GroupOnlineResponse)
	if err := c.Get("groups.online", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupRoles(params url.Values) (*GroupRolesResponse, error) {

	response := new(GroupRolesResponse)
	if err := c.Get("groups.roles", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupInfo(params url.Values) (*GroupInfoResponse, error) {

	response := new(GroupInfoResponse)
	if err := c.Get("groups.info", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupType(req *models.GroupTypeRequest) (*GroupTypeResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupTypeResponse)
	err = c.Post("groups.setType", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupHistory(params url.Values) (*GroupHistoryResponse, error) {

	response := new(GroupHistoryResponse)
	if err := c.Get("groups.history", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GroupLeave(req *models.GroupLeaveRequest) (*GroupLeaveResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupLeaveResponse)
	err = c.Post("groups.leave", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupAddAll(req *models.GroupAddAllRequest) (*GroupAddAllResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupAddAllResponse)
	err = c.Post("groups.addAll", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupArchive(req *models.GroupArchiveRequest) (*GroupArchiveResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupArchiveResponse)
	err = c.Post("groups.archive", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupUnArchive(req *models.GroupUnArchiveRequest) (*GroupUnArchiveResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupUnArchiveResponse)
	err = c.Post("groups.unarchive", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupDelete(req *models.GroupDeleteRequest) (*GroupDeleteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupDeleteResponse)
	err = c.Post("groups.delete", bytes.NewBuffer(body), response)
	return response, nil
}

func (c *Client) GroupCustomFields(req *models.GroupCustomFieldsRequest) (*GroupCustomFieldsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GroupCustomFieldsResponse)
	err = c.Post("groups.setCustomFields", bytes.NewBuffer(body), response)
	return response, nil
}
