package models

import "time"

type AgentInfoRequest struct {
	RID   string `json:"rid"`
	Token string `json:"token"`
}

type Agent struct {
	ID     string `json:"_id"`
	Emails []struct {
		Address  string `json:"address"`
		Verified bool   `json:"verified"`
	} `json:"emails"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type AgentDepartment struct {
	ID           string    `json:"_id"`
	AgentID      string    `json:"agentId"`
	DepartmentID string    `json:"departmentId"`
	UpdatedAt    time.Time `json:"_updatedAt"`
	Count        int       `json:"count"`
	Order        int       `json:"order"`
	Username     string    `json:"username"`
}

type Appearance struct {
	ID              string    `json:"_id"`
	UpdatedAt       time.Time `json:"_updatedAt"`
	Autocomplete    bool      `json:"autocomplete"`
	Blocked         bool      `json:"blocked"`
	CreatedAt       time.Time `json:"createdAt"`
	Enterprise      bool      `json:"enterprise"`
	Group           string    `json:"group"`
	Hidden          bool      `json:"hidden"`
	I18NDescription string    `json:"i18nDescription"`
	I18NLabel       string    `json:"i18nLabel"`
	PackageValue    string    `json:"packageValue"`
	Public          bool      `json:"public,omitempty"`
	Secret          bool      `json:"secret"`
	Section         string    `json:"section"`
	Sorter          int       `json:"sorter"`
	Ts              time.Time `json:"ts"`
	Type            string    `json:"type"`
	Value           string    `json:"value"`
	ValueSource     string    `json:"valueSource"`
	Multiline       bool      `json:"multiline,omitempty"`
	AllowedTypes    []string  `json:"allowedTypes,omitempty"`
	Editor          string    `json:"editor,omitempty"`
	PackageEditor   string    `json:"packageEditor,omitempty"`
	EnableQuery     string    `json:"enableQuery,omitempty"`
}

type Settings struct {
	RegistrationForm           bool   `json:"registrationForm"`
	AllowSwitchingDepartments  bool   `json:"allowSwitchingDepartments"`
	NameFieldRegistrationForm  bool   `json:"nameFieldRegistrationForm"`
	EmailFieldRegistrationForm bool   `json:"emailFieldRegistrationForm"`
	DisplayOfflineForm         bool   `json:"displayOfflineForm"`
	VideoCall                  bool   `json:"videoCall"`
	FileUpload                 bool   `json:"fileUpload"`
	Language                   string `json:"language"`
	Transcript                 bool   `json:"transcript"`
	HistoryMonitorType         string `json:"historyMonitorType"`
}

type Theme struct {
	Title        string        `json:"title"`
	Color        string        `json:"color"`
	OfflineTitle string        `json:"offlineTitle"`
	OfflineColor string        `json:"offlineColor"`
	ActionLinks  []ActionLinks `json:"actionLinks"`
}

type ActionLinks struct {
	Icon      string `json:"icon"`
	I18NLabel string `json:"i18nLabel"`
	MethodID  string `json:"method_id"`
	Params    string `json:"params"`
}

type LiveChatMessage struct {
	OfflineMessage              string `json:"offlineMessage"`
	OfflineSuccessMessage       string `json:"offlineSuccessMessage"`
	OfflineUnavailableMessage   string `json:"offlineUnavailableMessage"`
	ConversationFinishedMessage string `json:"conversationFinishedMessage"`
	TranscriptMessage           string `json:"transcriptMessage"`
}

type Survey struct {
	Items  []string `json:"items"`
	Values []string `json:"values"`
}

type Guest struct {
	GuestID  string `json:"_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Name     string `json:"name"`
}

type LiveChatRoomTransfer struct {
	RoomID   string `json:"_id"`
	ServedBy struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
	} `json:"servedBy"`
	Open bool `json:"open"`
}

type LiveChatConfig struct {
	Enabled  bool `json:"enabled"`
	Settings struct {
		RegistrationForm                 bool   `json:"registrationForm"`
		AllowSwitchingDepartments        bool   `json:"allowSwitchingDepartments"`
		NameFieldRegistrationForm        bool   `json:"nameFieldRegistrationForm"`
		EmailFieldRegistrationForm       bool   `json:"emailFieldRegistrationForm"`
		DisplayOfflineForm               bool   `json:"displayOfflineForm"`
		VideoCall                        bool   `json:"videoCall"`
		FileUpload                       bool   `json:"fileUpload"`
		Language                         string `json:"language"`
		Transcript                       bool   `json:"transcript"`
		HistoryMonitorType               string `json:"historyMonitorType"`
		ForceAcceptDataProcessingConsent bool   `json:"forceAcceptDataProcessingConsent"`
		ShowConnecting                   bool   `json:"showConnecting"`
		AgentHiddenInfo                  bool   `json:"agentHiddenInfo"`
	} `json:"settings"`
	Theme struct {
		Title        string `json:"title"`
		Color        string `json:"color"`
		OfflineTitle string `json:"offlineTitle"`
		OfflineColor string `json:"offlineColor"`
		ActionLinks  []struct {
			Icon      string `json:"icon"`
			I18NLabel string `json:"i18nLabel"`
			MethodID  string `json:"method_id"`
			Params    string `json:"params"`
		} `json:"actionLinks"`
	} `json:"theme"`
	Messages struct {
		OfflineMessage              string `json:"offlineMessage"`
		OfflineSuccessMessage       string `json:"offlineSuccessMessage"`
		OfflineUnavailableMessage   string `json:"offlineUnavailableMessage"`
		ConversationFinishedMessage string `json:"conversationFinishedMessage"`
		ConversationFinishedText    string `json:"conversationFinishedText"`
		TranscriptMessage           string `json:"transcriptMessage"`
		RegistrationFormMessage     string `json:"registrationFormMessage"`
		DataProcessingConsentText   string `json:"dataProcessingConsentText"`
	} `json:"messages"`
	Survey struct {
		Items  []string `json:"items"`
		Values []string `json:"values"`
	} `json:"survey"`
	Triggers []struct {
		ID      string `json:"_id"`
		Actions []struct {
			Name   string `json:"name"`
			Params struct {
				Sender string `json:"sender"`
				Msg    string `json:"msg"`
			} `json:"params"`
		} `json:"actions"`
		Conditions []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"conditions"`
		RunOnce bool `json:"runOnce"`
	} `json:"triggers"`
	Departments []struct {
		ID                 string `json:"_id"`
		Name               string `json:"name"`
		ShowOnRegistration bool   `json:"showOnRegistration"`
		ShowOnOfflineForm  bool   `json:"showOnOfflineForm"`
	} `json:"departments"`
	Resources struct {
		Sound  string        `json:"sound"`
		Emojis []interface{} `json:"emojis"`
	} `json:"resources"`
	Online bool `json:"online"`
	Guest  struct {
		ID            string `json:"_id"`
		Username      string `json:"username"`
		Name          string `json:"name"`
		Token         string `json:"token"`
		VisitorEmails []struct {
			Address string `json:"address"`
		} `json:"visitorEmails"`
	} `json:"guest"`
	Room struct {
		ID       string `json:"_id"`
		Open     bool   `json:"open"`
		ServedBy struct {
			ID       string    `json:"_id"`
			Username string    `json:"username"`
			Ts       time.Time `json:"ts"`
		} `json:"servedBy"`
	} `json:"room"`
	Agent struct {
		ID       string `json:"_id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Status   string `json:"status"`
	} `json:"agent"`
	CustomFields []interface{} `json:"customFields"`
}

type Department struct {
	ID                 string    `json:"id"`
	Enabled            bool      `json:"enabled"`
	ShowOnRegistration bool      `json:"showOnRegistration"`
	Email              string    `json:"email"`
	ShowOnOfflineForm  bool      `json:"showOnOfflineForm"`
	Name               string    `json:"name"`
	UpdatedAt          time.Time `json:"_updatedAt"`
	D                  string    `json:"_id,omitempty"`
	Description        string    `json:"description,omitempty"`
}

type Inquiry struct {
	ID      string    `json:"_id"`
	Rid     string    `json:"rid"`
	Message string    `json:"message"`
	Name    string    `json:"name"`
	Ts      time.Time `json:"ts"`
	Agents  []string  `json:"agents"`
	Status  string    `json:"status"`
	V       struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
		Token    string `json:"token"`
		Status   string `json:"status"`
	} `json:"v"`
	T         string    `json:"t"`
	UpdatedAt time.Time `json:"_updatedAt"`
}

type IntegrationSetting struct {
	ID                  string    `json:"_id"`
	UpdatedAt           time.Time `json:"_updatedAt"`
	Autocomplete        bool      `json:"autocomplete"`
	Blocked             bool      `json:"blocked"`
	CreatedAt           time.Time `json:"createdAt"`
	Group               string    `json:"group"`
	Hidden              bool      `json:"hidden"`
	I18NDescription     string    `json:"i18nDescription"`
	I18NLabel           string    `json:"i18nLabel"`
	PackageValue        string    `json:"packageValue"`
	Secret              bool      `json:"secret"`
	Section             string    `json:"section"`
	Sorter              int       `json:"sorter"`
	Ts                  time.Time `json:"ts"`
	Type                string    `json:"type"`
	Value               string    `json:"value"`
	ValueSource         string    `json:"valueSource"`
	MeteorSettingsValue string    `json:"meteorSettingsValue"`
}

type LiveChatOfficeHour struct {
	ID        string    `json:"_id"`
	Day       string    `json:"day"`
	Start     string    `json:"start"`
	Finish    string    `json:"finish"`
	Code      int       `json:"code"`
	Open      bool      `json:"open"`
	UpdatedAt time.Time `json:"_updatedAt"`
}

type LiveChatUsers struct {
	ID       string `json:"_id"`
	UserName string `json:"username"`
}

type LiveChatRoom struct {
	ID         string    `json:"_id"`
	Msgs       int       `json:"msgs"`
	UsersCount int       `json:"usersCount"`
	Lm         time.Time `json:"lm"`
	Fname      string    `json:"fname"`
	T          string    `json:"t"`
	Ts         time.Time `json:"ts"`
	V          struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
		Token    string `json:"token"`
		Status   string `json:"status"`
	} `json:"v"`
	Cl              bool      `json:"cl"`
	Open            bool      `json:"open"`
	WaitingResponse bool      `json:"waitingResponse"`
	UpdatedAt       time.Time `json:"_updatedAt"`
	ServedBy        struct {
		ID       string    `json:"_id"`
		Username string    `json:"username"`
		Ts       time.Time `json:"ts"`
	}
}

type LiveChatTrigger struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	RunOnce     bool   `json:"runOnce"`
	Conditions  []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"conditions"`
	Actions []struct {
		Name   string `json:"name"`
		Params struct {
			Sender string `json:"sender"`
			Msg    string `json:"msg"`
		} `json:"params"`
	} `json:"actions"`
	UpdatedAt time.Time `json:"_updatedAt"`
}

type LiveChatUser struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
}

type LiveChatVisitor struct {
	ID        string    `json:"_id"`
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"_updatedAt"`
	Token     string    `json:"token"`
	Phone     []struct {
		PhoneNumber string `json:"phoneNumber"`
	} `json:"phone"`
	VisitorEmails []struct {
		Address string `json:"address"`
	} `json:"visitorEmails"`
	Name         string `json:"name"`
	LivechatData struct {
		Address string `json:"address"`
	} `json:"livechatData"`
}

type LiveChatPage struct {
	ID  string    `json:"_id"`
	T   string    `json:"t"`
	Rid string    `json:"rid"`
	Ts  time.Time `json:"ts"`
	Msg string    `json:"msg"`
	U   struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
	} `json:"u"`
	Groupable  bool `json:"groupable"`
	Navigation struct {
		Page struct {
			Change   string `json:"change"`
			Title    string `json:"title"`
			Location struct {
				Href string `json:"href"`
			} `json:"location"`
		} `json:"page"`
		Token string `json:"token"`
	} `json:"navigation"`
	Hidden    bool      `json:"_hidden"`
	UpdatedAt time.Time `json:"_updatedAt"`
}

type LiveChatHistory struct {
	ID         string    `json:"_id"`
	Msgs       int       `json:"msgs"`
	UsersCount int       `json:"usersCount"`
	Lm         time.Time `json:"lm"`
	Fname      string    `json:"fname"`
	T          string    `json:"t"`
	Ts         time.Time `json:"ts"`
	V          struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
		Token    string `json:"token"`
		Status   string `json:"status"`
	} `json:"v"`
	Cl              bool      `json:"cl"`
	Open            bool      `json:"open"`
	WaitingResponse bool      `json:"waitingResponse"`
	UpdatedAt       time.Time `json:"_updatedAt"`
	ServedBy        struct {
		ID       string    `json:"_id"`
		Username string    `json:"username"`
		Ts       time.Time `json:"ts"`
	} `json:"servedBy"`
}
