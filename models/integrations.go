package models

import "time"

type Integration struct {
	ID        string    `json:"_id"`
	CreatedAt time.Time `json:"_createdAt"`
	CreatedBy CreatedBy `json:"_createdBy"`
	UpdatedAt time.Time `json:"_updatedAt"`
	CreateIntegration
	ImpersonateUser bool        `json:"impersonateUser"`
	ScriptCompiled  string      `json:"scriptCompiled"`
	ScriptError     interface{} `json:"scriptError"`
	UserId          string      `json:"userId"`
}

type CreatedBy struct {
	Username string `json:"username"`
	ID       string `json:"_id"`
}

type CreateIntegration struct {
	Type          string   `json:"type"`
	Name          string   `json:"name"`
	Event         string   `json:"event"`
	Enabled       bool     `json:"enabled"`
	Username      string   `json:"username"`
	Urls          []string `json:"urls"`
	ScriptEnabled bool     `json:"scriptEnabled"`
	Channel       string   `json:"channel"`
	TriggerWords  []string `json:"triggerWords"`
	Alias         string   `json:"alias"`
	Avater        string   `json:"avater"`
	Emoji         string   `json:"emoji"`
	Token         string   `json:"token"`
	Script        string   `json:"script"`
}

type PartIntegrationHistory struct {
	ID   string `json:"_id"`
	Type string `json:"type"`
	Step string `json:"step"`
	Integration
}

type IntegrationHistory struct {
	PartIntegrationHistory
	Event           string    `json:"event"`
	CreateAt        time.Time `json:"_createdAt"`
	UpdateAt        time.Time `json:"_updatedAt"`
	IntegrationData `json:"data"`
	IntegrationRoom `json:"room"`
	Url             string `json:"url"`
}

type IntegrationData struct {
	Token       string    `json:"token"`
	Bot         bool      `json:"bot"`
	TimeStamp   time.Time `json:"timestamp"`
	ChannelId   string    `json:"channel_id"`
	ChannelName string    `json:"channel_name"`
	UserId      string    `json:"user_id"`
	UserName    string    `json:"user_name"`
	User        User
}

type IntegrationRoom struct {
	ID          string        `json:"_id"`
	TimeStamp   time.Time     `json:"ts"`
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
	Lm           time.Time `json:"lm"`
	JitsiTimeout time.Time `json:"jitsiTimeout"`
}
