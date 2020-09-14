package models

import "time"

type CreateIMRequest struct {
	Username  string `json:"username"`
	Usernames string `json:"usernames"`
}

type CountersIMMessage struct {
	Joined       bool      `json:"joined"`
	Members      int       `json:"members"`
	Unreads      int       `json:"unreads"`
	UnreadsFrom  time.Time `json:"unreadsFrom"`
	Msgs         int       `json:"msgs"`
	Latest       time.Time `json:"latest"`
	UserMentions int       `json:"userMentions"`
}

type CloseIMRequest struct {
	RoomId string `json:"roomId"`
}

type FileIM struct {
	ID          string    `json:"_id"`
	Name        string    `json:"name"`
	Size        int       `json:"size"`
	Type        string    `json:"type"`
	RoomId      string    `json:"rid"`
	Description string    `json:"description"`
	Store       string    `json:"store"`
	Complete    bool      `json:"complete"`
	Uploading   bool      `json:"uploading"`
	Extension   string    `json:"extension"`
	Progress    int       `json:"progress"`
	User        User      `json:"user"`
	UpdatedAt   time.Time `json:"_updatedAt"`
	InstanceID  string    `json:"instanceId"`
	Etag        string    `json:"etag"`
	Path        string    `json:"path"`
	Token       string    `json:"token"`
	UploadedAt  time.Time `json:"uploadedAt"`
	URL         string    `json:"url"`
}

type MemberIM struct {
	ID        string `json:"_id"`
	Status    string `json:"status"`
	Name      string `json:"name"`
	UtcOffset int    `json:"utcOffset"`
	Username  string `json:"username"`
}

type ListIM struct {
	ID        string    `json:"_id"`
	UpdatedAt time.Time `json:"_updatedAt"`
	Type      string    `json:"t"`
	Msgs      int       `json:"msgs"`
	Ts        time.Time `json:"ts"`
	Lm        time.Time `json:"lm"`
	Topic     string    `json:"topic"`
}

type EveryoneListIM struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Type      string    `json:"t"`
	Usernames []string  `json:"usernames"`
	Msgs      int       `json:"msgs"`
	User      User      `json:"u"`
	Ts        time.Time `json:"ts"`
	Ro        bool      `json:"ro"`
	SysMes    bool      `json:"sysMes"`
	UpdatedAt time.Time `json:"_updatedAt"`
}
