package models

import "time"

type User struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	UserName     string `json:"username"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	TokenExpires int64  `json:"tokenExpires"`
}

type CreateUserRequest struct {
	Name         string            `json:"name"`
	Email        string            `json:"email"`
	Password     string            `json:"password"`
	Username     string            `json:"username"`
	Roles        []string          `json:"roles,omitempty"`
	CustomFields map[string]string `json:"customFields,omitempty"`
}

type RegisterUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"pass"`
	Username  string `json:"username"`
	SecretURL string `json:"secretURL"`
}

type UpdateUserRequest struct {
	UserID string `json:"userId"`
	Data   struct {
		Name         string            `json:"name"`
		Email        string            `json:"email"`
		Password     string            `json:"password"`
		Username     string            `json:"username"`
		CustomFields map[string]string `json:"customFields,omitempty"`
	} `json:"data"`
}

type UsersList struct {
	ID        string `json:"_id"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	Active    bool   `json:"active"`
	Name      string `json:"name"`
	Utcoffset bool   `json:"utcoffset"`
	Username  string `json:"username"`
}

type UserInfo struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	UtcOffset int    `json:"utcOffset"`
	Active    bool   `json:"active"`
	Type      string `json:"type"`
}

type UserResetAvatarRequest struct {
	UserId   string `json:"userId"`
	UserName string `json:"username"`
}

type Preferences struct {
	AutoImageLoad                         bool          `json:"autoImageLoad"`
	CollapseMediaByDefault                bool          `json:"collapseMediaByDefault"`
	ConvertASCIIEmoji                     bool          `json:"convertAsciiEmoji"`
	DesktopNotificationDuration           int           `json:"desktopNotificationDuration"`
	DesktopNotificationRequireInteraction bool          `json:"desktopNotificationRequireInteraction"`
	DesktopNotifications                  string        `json:"desktopNotifications"`
	DontAskAgainList                      []interface{} `json:"dontAskAgainList"`
	EmailNotificationMode                 string        `json:"emailNotificationMode"`
	EnableAutoAway                        bool          `json:"enableAutoAway"`
	HideAvatars                           bool          `json:"hideAvatars"`
	HideFlexTab                           bool          `json:"hideFlexTab"`
	HideRoles                             bool          `json:"hideRoles"`
	HideUsernames                         bool          `json:"hideUsernames"`
	Highlights                            []interface{} `json:"highlights"`
	IdleTimeLimit                         int           `json:"idleTimeLimit"`
	Language                              string        `json:"language"`
	MessageViewMode                       int           `json:"messageViewMode"`
	MobileNotifications                   string        `json:"mobileNotifications"`
	MuteFocusedConversations              bool          `json:"muteFocusedConversations"`
	NewMessageNotification                string        `json:"newMessageNotification"`
	NewRoomNotification                   string        `json:"newRoomNotification"`
	NotificationsSoundVolume              int           `json:"notificationsSoundVolume"`
	SaveMobileBandwidth                   bool          `json:"saveMobileBandwidth"`
	SendOnEnter                           string        `json:"sendOnEnter"`
	SidebarHideAvatar                     bool          `json:"sidebarHideAvatar"`
	SidebarShowDiscussion                 bool          `json:"sidebarShowDiscussion"`
	UnreadAlert                           bool          `json:"unreadAlert"`
	UseEmojis                             bool          `json:"useEmojis"`
}

type SetPreferrences struct {
	UserID string `json:"userId"`
	Data   struct {
		AutoImageLoad                         bool          `json:"autoImageLoad"`
		CollapseMediaByDefault                bool          `json:"collapseMediaByDefault"`
		ConvertASCIIEmoji                     bool          `json:"convertAsciiEmoji"`
		DesktopNotificationDuration           int           `json:"desktopNotificationDuration"`
		DesktopNotificationRequireInteraction bool          `json:"desktopNotificationRequireInteraction"`
		DesktopNotifications                  string        `json:"desktopNotifications"`
		EmailNotificationMode                 string        `json:"emailNotificationMode"`
		EnableAutoAway                        bool          `json:"enableAutoAway"`
		GroupByType                           bool          `json:"groupByType"`
		HideAvatars                           bool          `json:"hideAvatars"`
		HideRoles                             bool          `json:"hideRoles"`
		HideUsernames                         bool          `json:"hideUsernames"`
		Highlights                            []interface{} `json:"highlights,omitempty"`
		Language                              string        `json:"language"`
		MobileNotifications                   string        `json:"mobileNotifications"`
		MuteFocusedConversations              bool          `json:"muteFocusedConversations"`
		NewMessageNotification                string        `json:"newMessageNotification"`
		NewRoomNotification                   string        `json:"newRoomNotification"`
		NotificationsSoundVolume              int           `json:"notificationsSoundVolume"`
		RoomsListExhibitionMode               string        `json:"roomsListExhibitionMode"`
		SaveMobileBandwidth                   bool          `json:"saveMobileBandwidth"`
		RoomCounterSidebar                    bool          `json:"roomCounterSidebar"`
		SendOnEnter                           string        `json:"sendOnEnter"`
		SidebarHideAvatar                     bool          `json:"sidebarHideAvatar"`
		SidebarShowDiscussion                 bool          `json:"sidebarShowDiscussion"`
		SidebarSortby                         string        `json:"sidebarSortby"`
		SidebarViewMode                       string        `json:"sidebarViewMode"`
		SidebarShowFavorites                  bool          `json:"sidebarShowFavorites"`
		UnreadAlert                           bool          `json:"unreadAlert"`
		ViewMode                              int           `json:"viewMode"`
		UseEmojis                             bool          `json:"useEmojis"`
	} `json:"data,omitempty"`
}

type SetStatus struct {
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

type GeneratePersonalAccessToken struct {
	TokenName       string `json:"tokenName"`
	BypassTwoFactor bool   `json:"bypassTwoFactor,omitempty"`
}

type ExportOperation struct {
	UserID        string        `json:"userId"`
	RoomList      interface{}   `json:"roomList"`
	Status        string        `json:"status"`
	FileList      []interface{} `json:"fileList"`
	GeneratedFile interface{}   `json:"generatedFile"`
	FullExport    bool          `json:"fullExport"`
	UserData      struct {
		ID        string    `json:"_id"`
		CreatedAt time.Time `json:"createdAt"`
		Services  struct {
			Password struct {
				Bcrypt string `json:"bcrypt"`
				Reset  struct {
					Token  string    `json:"token"`
					Email  string    `json:"email"`
					When   time.Time `json:"when"`
					Reason string    `json:"reason"`
				} `json:"reset"`
			} `json:"password"`
			Email2Fa struct {
				Enabled   bool      `json:"enabled"`
				ChangedAt time.Time `json:"changedAt"`
			} `json:"email2fa"`
			Email struct {
				VerificationTokens []struct {
					Token   string    `json:"token"`
					Address string    `json:"address"`
					When    time.Time `json:"when"`
				} `json:"verificationTokens"`
			} `json:"email"`
			Resume struct {
				LoginTokens []struct {
					When          time.Time `json:"when,omitempty"`
					HashedToken   string    `json:"hashedToken"`
					Type          string    `json:"type,omitempty"`
					CreatedAt     time.Time `json:"createdAt,omitempty"`
					LastTokenPart string    `json:"lastTokenPart,omitempty"`
					Name          string    `json:"name,omitempty"`
				} `json:"loginTokens"`
			} `json:"resume"`
		} `json:"services"`
		Emails []struct {
			Address  string `json:"address"`
			Verified bool   `json:"verified"`
		} `json:"emails"`
		Type             string    `json:"type"`
		Status           string    `json:"status"`
		Active           bool      `json:"active"`
		UpdatedAt        time.Time `json:"_updatedAt"`
		Roles            []string  `json:"roles"`
		Name             string    `json:"name"`
		LastLogin        time.Time `json:"lastLogin"`
		StatusConnection string    `json:"statusConnection"`
		UtcOffset        int       `json:"utcOffset"`
		Username         string    `json:"username"`
		Settings         struct {
			Preferences struct {
				AutoImageLoad                         bool          `json:"autoImageLoad"`
				CollapseMediaByDefault                bool          `json:"collapseMediaByDefault"`
				ConvertASCIIEmoji                     bool          `json:"convertAsciiEmoji"`
				DesktopNotificationDuration           int           `json:"desktopNotificationDuration"`
				DesktopNotificationRequireInteraction bool          `json:"desktopNotificationRequireInteraction"`
				DesktopNotifications                  string        `json:"desktopNotifications"`
				DontAskAgainList                      []interface{} `json:"dontAskAgainList"`
				EmailNotificationMode                 string        `json:"emailNotificationMode"`
				EnableAutoAway                        bool          `json:"enableAutoAway"`
				HideAvatars                           bool          `json:"hideAvatars"`
				HideFlexTab                           bool          `json:"hideFlexTab"`
				HideRoles                             bool          `json:"hideRoles"`
				HideUsernames                         bool          `json:"hideUsernames"`
				Highlights                            []interface{} `json:"highlights"`
				IdleTimeLimit                         int           `json:"idleTimeLimit"`
				Language                              string        `json:"language"`
				MessageViewMode                       int           `json:"messageViewMode"`
				MobileNotifications                   string        `json:"mobileNotifications"`
				MuteFocusedConversations              bool          `json:"muteFocusedConversations"`
				NewMessageNotification                string        `json:"newMessageNotification"`
				NewRoomNotification                   string        `json:"newRoomNotification"`
				NotificationsSoundVolume              int           `json:"notificationsSoundVolume"`
				SaveMobileBandwidth                   bool          `json:"saveMobileBandwidth"`
				SendOnEnter                           string        `json:"sendOnEnter"`
				SidebarHideAvatar                     bool          `json:"sidebarHideAvatar"`
				SidebarShowDiscussion                 bool          `json:"sidebarShowDiscussion"`
				UnreadAlert                           bool          `json:"unreadAlert"`
				UseEmojis                             bool          `json:"useEmojis"`
			} `json:"preferences"`
		} `json:"settings"`
		Language      string `json:"language"`
		StatusText    string `json:"statusText"`
		StatusDefault string `json:"statusDefault"`
	} `json:"userData"`
	ID         string `json:"_id"`
	ExportPath string `json:"exportPath"`
	AssetsPath string `json:"assetsPath"`
}

type ForgotPassword struct {
	Email string `json:"email"`
}

type UpdateOwnBasicInfo struct {
	Data struct {
		Email           string `json:"email"`
		NewPassword     string `json:"newPassword"`
		CurrentPassword string `json:"currentPassword"`
		Username        string `json:"username"`
		Name            string `json:"name"`
	} `json:"data"`
}

type SetActiveStatus struct {
	ActiveStatus bool   `json:"activeStatus"`
	UserID       string `json:"userId"`
}

type UserDeactivateIdleRequest struct {
	DaysIdle int    `json:"daysIdle"`
	Role     string `json:"role"`
}

type UserDeleteRequest struct {
	UserID   string `json:"userId"`
	UserName string `json:"username"`
}

type UserDeleteOwnRequest struct {
	PassWord string `json:"password"`
}
