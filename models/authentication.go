package models

type AuthData struct {
	AuthToken string `json:"authToken"`
	UserID    string `json:"userId"`
	Me        Me     `json:"me"`
}

type Me struct {
	ID     string `json:"_id"`
	Name   string `json:"name"`
	Emails []struct {
		Address  string `json:"address"`
		Verified bool   `json:"verified"`
	} `json:"emails"`
	Status           string   `json:"status"`
	StatusConnection string   `json:"statusConnection"`
	Username         string   `json:"username"`
	UtcOffset        int      `json:"utcOffset"`
	Active           bool     `json:"active"`
	Roles            []string `json:"roles"`
	Settings         struct {
		Preferences struct {
			EnableAutoAway              bool   `json:"enableAutoAway"`
			IdleTimeoutLimit            int    `json:"idleTimeoutLimit"`
			DesktopNotificationDuration int    `json:"desktopNotificationDuration"`
			AudioNotifications          string `json:"audioNotifications"`
			DesktopNotifications        string `json:"desktopNotifications"`
			MobileNotifications         string `json:"mobileNotifications"`
			UnreadAlert                 bool   `json:"unreadAlert"`
			UseEmojis                   bool   `json:"useEmojis"`
			ConvertASCIIEmoji           bool   `json:"convertAsciiEmoji"`
			AutoImageLoad               bool   `json:"autoImageLoad"`
			SaveMobileBandwidth         bool   `json:"saveMobileBandwidth"`
			CollapseMediaByDefault      bool   `json:"collapseMediaByDefault"`
			HideUsernames               bool   `json:"hideUsernames"`
			HideRoles                   bool   `json:"hideRoles"`
			HideFlexTab                 bool   `json:"hideFlexTab"`
			HideAvatars                 bool   `json:"hideAvatars"`
			RoomsListExhibitionMode     string `json:"roomsListExhibitionMode"`
			SidebarViewMode             string `json:"sidebarViewMode"`
			SidebarHideAvatar           bool   `json:"sidebarHideAvatar"`
			SidebarShowUnread           bool   `json:"sidebarShowUnread"`
			SidebarShowFavorites        bool   `json:"sidebarShowFavorites"`
			SendOnEnter                 string `json:"sendOnEnter"`
			MessageViewMode             int    `json:"messageViewMode"`
			EmailNotificationMode       string `json:"emailNotificationMode"`
			RoomCounterSidebar          bool   `json:"roomCounterSidebar"`
			NewRoomNotification         string `json:"newRoomNotification"`
			NewMessageNotification      string `json:"newMessageNotification"`
			MuteFocusedConversations    bool   `json:"muteFocusedConversations"`
			NotificationsSoundVolume    int    `json:"notificationsSoundVolume"`
		} `json:"preferences"`
	} `json:"settings"`
	CustomFields struct {
		Twitter  string `json:"twitter"`
		Google   string `json:"google"`
		Facebook string `json:"facebook"`
	} `json:"customFields" omitempty`
	AvatarURL string `json:"avatarUrl"`
}
