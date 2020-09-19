package models

import "time"

type Channel struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Fname string `json:"fname,omitempty"`
	Type  string `json:"t"`
	Msgs  int    `json:"msgs"`

	ReadOnly  bool `json:"ro,omitempty"`
	SysMes    bool `json:"sysMes,omitempty"`
	Default   bool `json:"default"`
	Broadcast bool `json:"broadcast,omitempty"`

	Timestamp *time.Time `json:"ts,omitempty"`
	UpdatedAt *time.Time `json:"_updatedAt,omitempty"`

	User        *User    `json:"u,omitempty"`
	LastMessage *Message `json:"lastMessage,omitempty"`

	// Lm          interface{} `json:"lm"`
	// CustomFields struct {
	// } `json:"customFields,omitempty"`
}

type ChannelSubscription struct {
	ID          string   `json:"_id"`
	Alert       bool     `json:"alert"`
	Name        string   `json:"name"`
	DisplayName string   `json:"fname"`
	Open        bool     `json:"open"`
	RoomId      string   `json:"rid"`
	Type        string   `json:"c"`
	User        User     `json:"u"`
	Roles       []string `json:"roles"`
	Unread      float64  `json:"unread"`
}

type OnlineChannel struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
}

type ChannelCreateRequest struct {
	Name     string   `json:"name"`
	Members  []string `json:"members"`
	ReadOnly bool     `json:"readOnly"`
}

type ChannelAddAllRequest struct {
	RoomID          string `json:"roomId"`
	ActiveUsersOnly bool   `json:"activeUsersOnly"`
}

type ChannelAddLeaderRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type ChannelRemoveLeaderRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type ChannelAddOwnerRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type ChannelRemoveOwnerRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type ChannelAddModeratorRequest struct {
	RoomName string `json:"roomName"`
	RoomID   string `json:"roomId"`
	UserID   string `json:"userId"`
}

type ChannelRemoveModeratorRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type ChannelInviteRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type ChannelDescriptionRequest struct {
	RoomID      string `json:"roomId"`
	Description string `json:"description"`
}

type ChannelJoinCodeRequest struct {
	RoomID   string `json:"roomId"`
	JoinCode string `json:"joinCode"`
}

type ChannelTopicRequest struct {
	RoomID string `json:"roomId"`
	Topic  string `json:"topic"`
}

type ChannelReadOnlyRequest struct {
	RoomID   string `json:"roomId"`
	ReadOnly bool   `json:"readOnly"`
}

type ChannelTypeRequest struct {
	RoomID string `json:"roomId"`
	Type   string `json:"type"`
}

type ChannelAnnouncementRequest struct {
	RoomID       string `json:"roomId"`
	Announcement string `json:"announcement"`
}

type ChannelRenameRequest struct {
	RoomID string `json:"roomId"`
	Name   string `json:"name"`
}

type ChannelArchiveRequest struct {
	RoomID string `json:"roomId"`
}

type ChannelUnArchiveRequest struct {
	RoomID string `json:"roomId"`
}

type ChannelCloseRequest struct {
	RoomID string `json:"roomId"`
}

type ChannelDeleteRequest struct {
	RoomID   string `json:"roomId"`
	RoomName string `json:"roomName"`
}

type ChannelJoinRequest struct {
	RoomID   string `json:"roomId"`
	JoinCode string `json:"joinCode"`
}

type ChannelKickRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type ChannelDefaultRequest struct {
	RoomID  string `json:"roomId"`
	Default bool   `json:"default"`
}

type ChannelOpenRequest struct {
	RoomID string `json:"roomId"`
}

type ChannelCustomFieldsRequest struct {
	RoomID       string `json:"roomId"`
	RoomName     string `json:"roomName"`
	CustomFields struct {
		Organization string `json:"organization"`
	} `json:"customFields"`
}
