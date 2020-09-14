package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

type UserGetAvatarResponse struct {
	Status
}

type UserDeleteOwnResponse struct {
	Status
}

type UserDeleteResponse struct {
	Status
}

type UserDeactivateIdleResponse struct {
	Status
	Count int `json:"count"`
}

type logoutResponse struct {
	Status
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

type LogonResponse struct {
	Status
	Data struct {
		Token  string `json:"authToken"`
		UserID string `json:"userID"`
	} `json:"data"`
}

type CreateUserResponse struct {
	Status
	User struct {
		ID        string    `json:"_id"`
		CreatedAt time.Time `json:"createdAt"`
		Services  struct {
			Password struct {
				Bcrypt string `json:"bcrypt"`
			} `json:"password"`
		} `json:"services"`
		Username string `json:"username"`
		Emails   []struct {
			Address  string `json:"address"`
			Verified bool   `json:"verified"`
		} `json:"emails"`
		Type         string            `json:"type"`
		Status       string            `json:"status"`
		Active       bool              `json:"active"`
		Roles        []string          `json:"roles"`
		UpdatedAt    time.Time         `json:"_updatedAt"`
		Name         string            `json:"name"`
		CustomFields map[string]string `json:"customFields"`
	} `json:"user"`
}

type RegisterUserResponse struct {
	Status
	User struct {
		ID        string `json:"_id"`
		Type      string `json:"type"`
		Status    string `json:"status"`
		Active    bool   `json:"active"`
		Name      string `json:"name"`
		UtcOffset int    `json:"utcOffset"`
		Username  string `json:"username"`
	} `json:"user"`
	Success bool `json:"success"`
}

type UsersListResponse struct {
	Status
	UsersList []models.UsersList `json:"users"`
}

type UserResetAvatarResponse struct {
	Status
}

type UsersInfoResponse struct {
	Status
	UserInfo models.UserInfo `json:"user"`
}

type UserPresenceResponse struct {
	Status
	// UserPresence models.UserPresence
	Presence string `json:"presence"`
}

type PresenceResponse struct {
	Status
	Users []struct {
		ID         string `json:"_id"`
		Status     string `json:"status"`
		Name       string `json:"name"`
		UtcOffset  int    `json:"utcOffset,omitempty"`
		Username   string `json:"username"`
		StatusText string `json:"statusText,omitempty"`
	} `json:"users"`
	Full bool `json:"full"`
}

type GetStatusResponse struct {
	Status
	ID               string `json:"_id"`
	ConnectionStatus string `json:"connectionStatus"`
}

type GetUsernameSuggestionResponse struct {
	Status
	Result string `json:"result"`
}

type GetPreferencesResponse struct {
	Status
	Preferences models.Preferences `json:"preferences"`
}

type SetStatusResponse struct {
	Status
}

type SetPreferencesResponse struct {
	Status
	User struct {
		ID       string `json:"_id"`
		Settings struct {
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
	} `json:"user"`
}

type GeneratePersonalAccessTokenResponse struct {
	Status
	Token string `json:"token"`
}

type GetPersonalAccessTokensResponse struct {
	Status
	Tokens []struct {
		Name            string    `json:"name"`
		CreatedAt       time.Time `json:"createdAt"`
		LastTokenPart   string    `json:"lastTokenPart"`
		BypassTwoFactor bool      `json:"bypassTwoFactor,omitempty"`
	} `json:"tokens"`
}

type RemoveGeneratePersonalAccessTokenResponse struct {
	Status
}

type RequestDataDownloadResponse struct {
	Status
	Requested       bool                   `json:"requested"`
	ExportOperation models.ExportOperation `json:"exportOperation"`
}

type ForgotPasswordResponse struct {
	Status
}

type SetActiveStatusResponse struct {
	Status
	User struct {
		ID     string `json:"_id"`
		Active bool   `json:"active"`
	} `json:"user"`
}

func (c *Client) UsersList(params url.Values) (*UsersListResponse, error) {
	response := new(UsersListResponse)
	if err := c.Get("users.list", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UserInfo(params url.Values) (*UsersInfoResponse, error) {
	response := new(UsersInfoResponse)
	if err := c.Get("users.info", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Presence(params url.Values) (*PresenceResponse, error) {
	response := new(PresenceResponse)
	if err := c.Get("users.presence", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetPresence(params url.Values) (*UserPresenceResponse, error) {
	response := new(UserPresenceResponse)
	if err := c.Get("users.getPresence", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetStatus(params url.Values) (*GetStatusResponse, error) {
	response := new(GetStatusResponse)
	if err := c.Get("users.getStatus", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetUsernameSuggestion() (*GetUsernameSuggestionResponse, error) {
	response := new(GetUsernameSuggestionResponse)
	if err := c.Get("users.getUsernameSuggestion", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetPreferences() (*GetPreferencesResponse, error) {
	response := new(GetPreferencesResponse)
	if err := c.Get("users.getPreferences", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SetPreferrences(req *models.SetPreferrences) (*SetPreferencesResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(SetPreferencesResponse)
	err = c.Post("users.setPreferences", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) SetStatus(req *models.SetStatus) (*SetStatusResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(SetStatusResponse)
	err = c.Post("users.setStatus", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) GeneratePersonalAccessToken(req *models.GeneratePersonalAccessToken) (*GeneratePersonalAccessTokenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	fmt.Println("req", req)
	response := new(GeneratePersonalAccessTokenResponse)
	err = c.Post("users.generatePersonalAccessToken", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) GetPersonalAccessTokens() (*GetPersonalAccessTokensResponse, error) {
	response := new(GetPersonalAccessTokensResponse)
	if err := c.Get("users.getPersonalAccessTokens", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) RegeneratePersonalAccessToken(req *models.GeneratePersonalAccessToken) (*GeneratePersonalAccessTokenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(GeneratePersonalAccessTokenResponse)
	err = c.Post("users.regeneratePersonalAccessToken", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) RemovePersonalAccessToken(req *models.GeneratePersonalAccessToken) (*RemoveGeneratePersonalAccessTokenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(RemoveGeneratePersonalAccessTokenResponse)
	err = c.Post("users.removePersonalAccessToken", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) RequestDataDownload(params url.Values) (*RequestDataDownloadResponse, error) {
	response := new(RequestDataDownloadResponse)
	if err := c.Get("users.requestDataDownload", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ForgotPassword(req *models.ForgotPassword) (*ForgotPasswordResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(ForgotPasswordResponse)
	err = c.Post("users.forgotPassword", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) UpdateOwnBasicInfo(req *models.UpdateOwnBasicInfo) (*Status, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(Status)
	err = c.Post("users.updateOwnBasicInfo", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) SetActiveStatus(req *models.SetActiveStatus) (*SetActiveStatusResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(SetActiveStatusResponse)
	err = c.Post("users.setActiveStatus", bytes.NewBuffer(body), response)
	return response, err
}

// Login a user. The Email and the Password are mandatory. The auth token of the user is stored in the Client instance.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/login
func (c *Client) Login(credentials *models.UserCredentials) (*LogonResponse, error) {
	if c.auth != nil {
		return nil, nil
	}

	if credentials.ID != "" && credentials.Token != "" {
		c.auth = &authInfo{id: credentials.ID, token: credentials.Token}
		return nil, nil
	}

	response := new(LogonResponse)

	var data url.Values
	if credentials.Email != "" {
		data = url.Values{"user": {credentials.Email}, "password": {credentials.Password}}
	} else {
		data = url.Values{"user": {credentials.Username}, "password": {credentials.Password}}
	}

	if err := c.PostForm("login", data, response); err != nil {
		return nil, err
	}

	c.auth = &authInfo{id: response.Data.UserID, token: response.Data.Token}
	credentials.ID, credentials.Token = response.Data.UserID, response.Data.Token
	return response, nil
}

// CreateToken creates an access token for a user
//
// https://rocket.chat/docs/developer-guides/rest-api/users/createtoken/
func (c *Client) CreateToken(userID, username string) (*models.UserCredentials, error) {
	response := new(LogonResponse)
	data := url.Values{"userId": {userID}, "username": {username}}
	if err := c.PostForm("users.createToken", data, response); err != nil {
		return nil, err
	}
	credentials := &models.UserCredentials{}
	credentials.ID, credentials.Token = response.Data.UserID, response.Data.Token
	return credentials, nil
}

// Logout a user. The function returns the response message of the server.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/logout
func (c *Client) Logout() (string, error) {

	if c.auth == nil {
		return "Was not logged in", nil
	}

	response := new(logoutResponse)
	if err := c.Get("logout", nil, response); err != nil {
		return "", err
	}

	return response.Data.Message, nil
}

//
func (c *Client) RegisterUser(req *models.RegisterUserRequest) (*RegisterUserResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(RegisterUserResponse)
	err = c.Post("users.register", bytes.NewBuffer(body), response)
	return response, err
}

// CreateUser being logged in with a user that has permission to do so.
//
// https://rocket.chat/docs/developer-guides/rest-api/users/create
func (c *Client) CreateUser(req *models.CreateUserRequest) (*CreateUserResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(CreateUserResponse)
	err = c.Post("users.create", bytes.NewBuffer(body), response)
	return response, err
}

// UpdateUser updates a user's data being logged in with a user that has permission to do so.
//
// https://rocket.chat/docs/developer-guides/rest-api/users/update/
func (c *Client) UpdateUser(req *models.UpdateUserRequest) (*CreateUserResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(CreateUserResponse)
	err = c.Post("users.update", bytes.NewBuffer(body), response)
	return response, err
}

// SetUserAvatar updates a user's avatar being logged in with a user that has permission to do so.
// Currently only passing an URL is possible.
//
// https://rocket.chat/docs/developer-guides/rest-api/users/setavatar/
func (c *Client) SetUserAvatar(userID, username, avatarURL string) (*Status, error) {
	body := fmt.Sprintf(`{ "userId": "%s","username": "%s","avatarUrl":"%s"}`, userID, username, avatarURL)
	response := new(Status)
	err := c.Post("users.setAvatar", bytes.NewBufferString(body), response)
	return response, err
}

func (c *Client) UserResetAvatar(req *models.UserResetAvatarRequest) (*UserResetAvatarResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(UserResetAvatarResponse)
	err = c.Post("users.resetAvatar", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) UserDeactivateIdle(req *models.UserDeactivateIdleRequest) (*UserDeactivateIdleResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(UserDeactivateIdleResponse)
	err = c.Post("users.deactivateIdle", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) UserDelete(req *models.UserDeleteRequest) (*UserDeleteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(UserDeleteResponse)
	err = c.Post("users.delete", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) UserDeleteOwn(req *models.UserDeleteOwnRequest) (*UserDeleteOwnResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(UserDeleteOwnResponse)
	err = c.Post("users.deleteOwnAccount", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) UserGetAvatar(params url.Values) (*UserGetAvatarResponse, error) {
	response := new(UserGetAvatarResponse)
	if err := c.Get("users.users.getAvatar", params, response); err != nil {
		return nil, err
	}

	return response, nil
}
