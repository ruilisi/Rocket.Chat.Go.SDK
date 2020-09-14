package models

import "time"

type ChatDeleteRequest struct {
	RoomID string `json:"roomId"`
	MsgID  string `json:"msgId"`
	AsUser bool   `json:"asUser"`
}

type ChatUpdateRequest struct {
	RoomID string `json:"roomId"`
	MsgID  string `json:"msgId"`
	Text   string `json:"text"`
}
type Status struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`

	Status  string `json:"status"`
	Message string `json:"message"`
}

type CHATU struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
type CHATMessage struct {
	ID          string        `json:"_id"`
	Alias       string        `json:"alias"`
	Msg         string        `json:"msg"`
	Attachments []interface{} `json:"attachments"`
	ParseUrls   bool          `json:"parseUrls"`
	Groupable   bool          `json:"groupable"`
	Ts          time.Time     `json:"ts"`
	U           CHATU         `json:"u"`
	Rid         string        `json:"rid"`
	UpdatedAt   time.Time     `json:"_updatedAt"`
	Mentions    []interface{} `json:"mentions"`
	Channels    []interface{} `json:"channels"`
}

type ChatIgnoreRequest struct {
	Rid    string `json:"rid"`
	UserId string `json:"userId"`
	Ignore bool   `json:"ignore"`
}

type ChatStarRequest struct {
	MessageId string `json:"messageId"`
}

type ChatPinRequest struct {
	MessageId string `json:"messageId"`
}

type ChatFollowRequest struct {
	MessageId string `json:"mid"`
}

type ChatUnPinRequest struct {
	MessageId string `json:"messageId"`
}

type ChatUnStarRequest struct {
	MessageId string `json:"messageId"`
}

type ChatUnFollowRequest struct {
	MessageId string `json:"mid"`
}

type ChatSendMessageRequest struct {
	Message struct {
		Rid string `json:"rid"`
		Msg string `json:"msg"`
		//	Alias       string `json:"alias"`
		Emoji string `json:"emoji"`
		//	Avatar      string `json:"avatar"`
		Attachments []struct {
			Color             string    `json:"color"`
			Text              string    `json:"text"`
			Ts                time.Time `json:"ts"`
			ThumbURL          string    `json:"thumb_url"`
			MessageLink       string    `json:"message_link"`
			Collapsed         bool      `json:"collapsed"`
			AuthorName        string    `json:"author_name"`
			AuthorLink        string    `json:"author_link"`
			AuthorIcon        string    `json:"author_icon"`
			Title             string    `json:"title"`
			TitleLink         string    `json:"title_link"`
			TitleLinkDownload bool      `json:"title_link_download"`
			ImageURL          string    `json:"image_url"`
			AudioURL          string    `json:"audio_url"`
			VideoURL          string    `json:"video_url"`
			Fields            []struct {
				Short bool   `json:"short"`
				Title string `json:"title"`
				Value string `json:"value"`
			} `json:"fields"`
		} `json:"attachments"`
	} `json:"message"`
}

type ChatReactRequest struct {
	MessageID   string `json:"messageId"`
	Emoji       string `json:"emoji"`
	ShouldReact bool   `json:"shouldReact"`
}

type ChatReportRequest struct {
	MessageID   string `json:"messageId"`
	Description string `json:"description"`
}
