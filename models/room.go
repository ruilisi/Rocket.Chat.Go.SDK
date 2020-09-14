package models

import "time"

type Room struct {
	ID         string        `json:"_id"`
	Ts         time.Time     `json:"ts"`
	T          string        `json:"t"`
	Name       string        `json:"name"`
	Usernames  []interface{} `json:"usernames"`
	Msgs       int           `json:"msgs"`
	UsersCount int           `json:"usersCount"`
	Default    bool          `json:"default"`
	Favorite   bool          `json:"favorite"`
	UpdatedAt  time.Time     `json:"_updatedAt"`
	Lm         time.Time     `json:"lm"`
}

type Notifications struct {
	DesktopNotifications        string `json:"desktopNotifications,omitempty"`
	DisableNotifications        string `json:"disableNotifications,omitempty"`
	EmailNotifications          string `json:"emailNotifications,omitempty"`
	AudioNotificationValue      string `json:"audioNotificationValue,omitempty"`
	DesktopNotificationDuration string `json:"desktopNotificationDuration,omitempty"`
	AudioNotifications          string `json:"audioNotifications,omitempty"`
	UnreadAlert                 string `json:"unreadAlert,omitempty"`
	HideUnreadStatus            string `json:"hideUnreadStatus,omitempty"`
	MobilePushNotifications     string `json:"mobilePushNotifications,omitempty"`
}

type Upload struct {
	File        string `json:"file"`
	Msg         string `json:"msg,omitempty"`
	Description string `json:"description,omitempty"`
	Tmid        string `json:"tmid,omitempty"`
}

type RoomCleanHistoryRequest struct {
	RoomID        string    `json:"roomId"`
	Latest        time.Time `json:"latest"`
	Oldest        time.Time `json:"oldest"`
	Inclusive     bool      `json:"inclusive"`
	ExcludePinned bool      `json:"excludePinned"`
	FilesOnly     bool      `json:"filesOnly"`
	Users         []string  `json:"users"`
}
