package models

import "time"

type CreateGroupRequest struct {
	Name     string   `json:"name"`
	Members  []string `json:"members"`
	ReadOnly bool     `json:"readOnly"`
}

type GroupAddLeaderRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}
type GroupRemoveLeaderRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type GroupList struct {
	Groups []struct {
		ID       string            `json:"_id"`
		Name     string            `json:"name"`
		T        string            `json:"t"`
		Msgs     int               `json:"msgs"`
		U        map[string]string `json:"u"`
		Ts       time.Time         `json:"ts"`
		Ro       bool              `json:"ro"`
		SysMes   bool              `json:"sysMes"`
		UpdateAt time.Time         `json:"_updatedAt"`
	} `json:"groups"`
}

type GroupAddModeratorRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type GroupRemoveModeratorRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type GroupAddOwnerRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type GroupRemoveOwnerRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type GroupRenameRequest struct {
	RoomID string `json:"roomId"`
	Name   string `json:"name"`
}

type GroupCloseRequest struct {
	RoomID string `json:"roomId"`
}

type GroupOpenRequest struct {
	RoomID string `json:"roomId"`
}

type GroupInviteRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type GroupKickRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type GroupAnnouncementRequest struct {
	RoomID       string `json:"roomId"`
	Announcement string `json:"announcement"`
}

type GroupDescriptionRequest struct {
	RoomID      string `json:"roomId"`
	Description string `json:"description"`
}

type GroupReadOnlyRequest struct {
	RoomID   string `json:"roomId"`
	ReadOnly bool   `json:"readOnly"`
}

type GroupTopicRequest struct {
	RoomID string `json:"roomId"`
	Topic  string `json:"topic"`
}

type GroupTypeRequest struct {
	RoomID string `json:"roomId"`
	Type   string `json:"type"`
}

type GroupLeaveRequest struct {
	RoomID string `json:"roomId"`
}

type GroupAddAllRequest struct {
	RoomID          string `json:"roomId"`
	ActiveUsersOnly bool   `json:"activeUsersOnly"`
}

type GroupArchiveRequest struct {
	RoomID string `json:"roomId"`
}

type GroupUnArchiveRequest struct {
	RoomID string `json:"roomId"`
}

type GroupDeleteRequest struct {
	RoomID   string `json:"roomId"`
	RoomName string `json:"roomName"`
}

type GroupCustomFieldsRequest struct {
	RoomID       string `json:"roomId"`
	RoomName     string `json:"roomName"`
	CustomFields struct {
		Organization string `json:"organization"`
	} `json:"customFields"`
}
