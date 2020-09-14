package rest

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"

	"github.com/ruilisi/Rocket.Chat.Go.SDK/models"
)

// type AgentInfoResponse struct {
// Agent models.Agent
// Status
// }

type AgentInfoResponse struct {
	Status
	Agent struct {
		ID       string `json:"_id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Status   string `json:"status"`
	} `json:"agent"`
}

type AgentDepartmentResponse struct {
	Departments []struct {
		ID           string    `json:"_id"`
		AgentID      string    `json:"agentId"`
		DepartmentID string    `json:"departmentId"`
		UpdatedAt    time.Time `json:"_updatedAt"`
		Count        int       `json:"count"`
		Order        int       `json:"order"`
		Username     string    `json:"username"`
	} `json:"departments"`
	Status
}

// type AgentDepartmentResponse struct {
// Departments []models.AgentDepartment
// Status
// }

type AppearanceResponse struct {
	Appearance []models.Appearance
	Status
}

type LiveChatConfigResponse struct {
	Config models.LiveChatConfig `json:"config"`
	Status
}

type LiveChatCustomFieldRequest struct {
	Token     string `json:"token"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Overwrite bool   `json:"overwrite"`
}

type LiveChatCustomFields struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Overwrite bool   `json:"overwrite"`
}

type LiveChatCustomFieldResponse struct {
	Field struct {
		Key       string `json:"key"`
		Value     string `json:"value"`
		Overwrite bool   `json:"overwrite"`
	} `json:"field"`
	Status
}

type LiveChatCustomFieldsRequest struct {
	Token        string         `json:"token"`
	CustomFields []CustomFields `json:"customFields"`
}

type LiveChatCustomFieldsResponse struct {
	Fields []struct {
		Key       string `json:"Key"`
		Value     string `json:"value"`
		Overwrite bool   `json:"overwrite"`
	} `json:"fields"`
	Status
}

type LiveChatCustom_FieldsResponse struct {
	CustomFields []struct {
		ID         string    `json:"_id"`
		Label      string    `json:"label"`
		Scope      string    `json:"scope"`
		Visibility string    `json:"visibility"`
		UpdatedAt  time.Time `json:"_updatedAt"`
	} `json:"customFields"`
	base
}

type LiveChatCustom_FieldResponse struct {
	CustomField struct {
		ID         string    `json:"_id"`
		Label      string    `json:"label"`
		Scope      string    `json:"scope"`
		Visibility string    `json:"visibility"`
		UpdatedAt  time.Time `json:"_updatedAt"`
	} `json:"customField"`
	Status
}

type LiveChatUpdateDepartmentResponse struct {
	Department models.Department        `json:"department"`
	Agents     []models.AgentDepartment `json:"agents"`
	Status
}

type LiveChatDepartmentRequest struct {
	Department models.Department        `json:"department"`
	Agents     []models.AgentDepartment `json:"agents"`
}

type LiveChatResgisterDepartmentRequest struct {
	Department models.Department        `json:"department"`
	Agents     []models.AgentDepartment `json:"agents,omitempty"`
}

type LiveChatUpdateDepartmentRequest struct {
	Department models.Department        `json:"department"`
	Agents     []models.AgentDepartment `json:"agents,omitempty"`
}

type LiveChatDepartmentsResponse struct {
	Departments []models.Department `json:"departments"`
	Status
}

type LiveChatQueryDepartmentResponse struct {
	Departments models.Department        `json:"department"`
	Agents      []models.AgentDepartment `json:"agents,omitempty"`
	Status
}

type LiveChatRegisterDepartmentResponse struct {
	Departments []models.Department `json:"departments"`
	Status
}

type LiveChatRemoveDepartmentResponse struct {
	Status
}

type LiveChatListInquiriesResponse struct {
	Inquiries []models.Inquiry
	base
}

type LiveChatTakeInquiriesRequest struct {
	InquiryId string `json:"inquiryId"`
	UserId    string `json:"userId"`
}

type LiveChatTakeInquiriesResponse struct {
	Inquiries []models.Inquiry
	Status
}
type LiveChatGetOneInquiriesRequest struct {
	RoomId string `json:"roomId"`
}
type LiveChatGetOneInquiriesResponse struct {
	Inquiry models.Inquiry
	Status
}

type LiveChatIntegraionSettingResponse struct {
	Settings []models.IntegrationSetting
	Status
}

type LiveChatMessageRequest struct {
	Token string `json:"token"`
	Rid   string `json:"rid"`
	Msg   string `json:"msg"`
	Id    string `json:"_id,omitempty"`
	Agent string `json:"agent,omitempty"`
}

type LiveChatMessageResponse struct {
	Message struct {
		ID  string `json:"_id"`
		Msg string `json:"msg"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		Ts time.Time `json:"ts"`
	} `json:"message"`
	Status
}

type LiveChatUpdateMessageRequest struct {
	Token string `json:"token"`
	Rid   string `json:"rid"`
	Msg   string `json:"msg"`
}

type LiveChatDeleteMessageRequest struct {
	Token string `json:"token"`
	Rid   string `json:"rid"`
}

type LiveChatUpdateMessageResponse struct {
	Message struct {
		ID  string `json:"_id"`
		Msg string `json:"msg"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		Ts time.Time `json:"ts"`
	} `json:"message"`
	Status
}

type LiveChatDeleteMessageRequset struct {
	ID string `json:"_id"`
}

type LiveChatDeleteMessageResponse struct {
	Message struct {
		ID  string `json:"_id"`
		Msg string `json:"msg"`
		U   struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Name     string `json:"name"`
		} `json:"u"`
		Ts time.Time `json:"ts"`
	} `json:"message"`
	Status
}

type LiveChatLoadHistoryRequest struct {
	Rid   string    `json:"rid`
	Token string    `json:"token"`
	Ts    time.Time `json:"ts"`
	End   time.Time `json:"end"`
	Limit int       `json:"limit"`
}

type LiveChatLoadHistoryResponse struct {
	Messages        []models.Message `json:"messages"`
	UnreadNotLoaded int              `json:"unreadNotLoaded"`
	Status
}

type LiveChatOfflineSendRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type LiveChatOfflineSendResponse struct {
	Message string `json:"message"`
	Status
}

type LiveChatOfficeHourResponse struct {
	OfficeHours []models.LiveChatOfficeHour
	Status
}

type LiveChatQueueResponse struct {
	base
	Queue []struct {
		Chats int `json:"chats"`
		User  struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
			Status   string `json:"status"`
		} `json:"user"`
		Department struct {
			ID   interface{} `json:"_id"`
			Name interface{} `json:"name"`
		} `json:"department"`
	} `json:"queue"`
}

type LiveChatRoomResponse struct {
	Room    models.LiveChatRoom `json:"room"`
	NewRoom bool                `json:"newRoom"`
	Status
}

type LiveChatRoomCloseRequest struct {
	Rid   string `json:"rid"`
	Token string `json:"token"`
}

type LiveChatRoomCloseResponse struct {
	Rid     string `json:"rid"`
	Comment string `json:"comment"`
	Status
}

type LiveChatRoomTransferRequest struct {
	Rid        string `json:"rid"`
	Token      string `json:"token"`
	Department string `json:"department"`
}

type LiveChatRoomTransferResponse struct {
	Room models.LiveChatRoomTransfer `json:"room"`
	Status
}

type LiveChatRoomForwardRequest struct {
	RoomId       string `json:"roomId"`
	UserId       string `json:"userId"`
	DepartmentId string `json:"departmentId"`
}

type LiveChatRoomForwardResponse struct {
	Status
}

type LiveChatRoomSurveyRequest struct {
	Rid   string `json:"rid"`
	Token string `json:"token"`
	Data  []Data `json:"data"`
}
type Data struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// type LiveChatRoomSurveyRequest struct {
// Rid   string             `json:"rid"`
// Token string             `json:"token"`
// Data  []LiveChatRoomData `json:"data"`
// }
//
// type LiveChatRoomData struct {
// Name  string `json:"name"`
// Value string `json:"value"`
// }

type LiveChatRoomSurveyResponse struct {
	Rid  string `json:"rid"`
	Data struct {
		Satisfaction       string `json:"satisfaction"`
		AgentResposiveness string `json:"agentResposiveness"`
	} `json:"data"`
	Status
}

type LiveChatRoomsResponse struct {
	Rooms []models.LiveChatRoom
	base
}

type LiveChatTranscriptRequest struct {
	Rid   string `json:"rid"`
	Token string `json:"token"`
	Email string `json:"email"`
}

type LiveChatTranscriptResponse struct {
	Message string `json:"message"`
	Status
}

type LiveChatTriggersResponse struct {
	Triggers []models.LiveChatTrigger `json:"triggers"`
	base
}

type LiveChatTriggerResponse struct {
	Trigger models.LiveChatTrigger `json:"trigger"`
	Status
}

type LiveChatGetUsersResponse struct {
	Users []models.LiveChatUsers `json:"users"`
	Status
}

type LiveChatRegisterUserRequset struct {
	// Type     string `json:"type"`
	Username string `json:"username"`
}

type LiveChatRegisterUserResponse struct {
	User models.LiveChatUser `json:"user"`
	Status
}

type LiveChatGetInfoUserResponse struct {
	User models.LiveChatUser `json:"user"`
	Status
}

type LiveChatRemoveUserResponse struct {
	Status
}

type LiveChatVideoCallResponse struct {
	VideoCall struct {
		Rid      string    `json:"rid"`
		Domain   string    `json:"domain"`
		Provider string    `json:"provider"`
		Room     string    `json:"room"`
		Timeout  time.Time `json:"timeout"`
	} `json:"videoCall"`
	Status
}

type LiveChatVisitorRetrieveResponse struct {
	Visitor models.LiveChatVisitor `json:"visitor"`
	Status
}

type LiveChatVisitorRegisterRequest struct {
	Visitor struct {
		Name         string         `json:"name"`
		Email        string         `json:"email"`
		Token        string         `json:"token"`
		Phone        string         `json:"phone"`
		CustomFields []CustomFields `json:"visitor"`
		// CustomFields []struct {
		// Key       string `json:"key"`
		// Value     string `json:"value"`
		// Overwrite bool   `json:"overwrite"`
		// } `json:"customFields"`
	} `json:"visitor"`
}

type CustomFields struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Overwrite bool   `json:"overwrite"`
}

type LiveChatVisitorRegisterResponse struct {
	Visitor models.LiveChatVisitor `json:"visitor"`
	Status
}

// type PageInfo struct{
//
// }
type LiveChatVisitorNavigationHistoryRequest struct {
	Token    string `json:"token"`
	Rid      string `json:"rid"`
	PageInfo struct {
		Change   string `json:"change"`
		Title    string `json:"title"`
		Location struct {
			Href string `json:"href"`
		} `json:"location"`
	} `json:"pageInfo"`
}

type LiveChatVisitorNavigationHistoryResponse struct {
	Page struct {
		Msg        string `json:"msg"`
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
	} `json:"page"`
	Status
}

type LiveChatVisitorInfoResponse struct {
	Visitor struct {
		ID        string    `json:"_id"`
		Username  string    `json:"username"`
		UpdatedAt time.Time `json:"_updatedAt"`
		Token     string    `json:"token"`
	} `json:"visitor"`
	Status
}

type LiveChatVisitorNavigationHistoryRetrieveResponse struct {
	Pages []models.LiveChatPage `json:"pages"`
	Status
}

// type LiveChatVisitorChatHistoryResponse struct {
// Histories []models.LiveChatHistory
// base
// }

type LiveChatVisitorChatHistoryResponse struct {
	base
	History []struct {
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
		} `json:"v,omitempty"`
		Cl              bool      `json:"cl"`
		Open            bool      `json:"open,omitempty"`
		WaitingResponse bool      `json:"waitingResponse,omitempty"`
		UpdatedAt       time.Time `json:"_updatedAt"`
		ServedBy        struct {
			ID       string    `json:"_id"`
			Username string    `json:"username"`
			Ts       time.Time `json:"ts"`
		} `json:"servedBy"`
		LastMessage struct {
			ID  string    `json:"_id"`
			Rid string    `json:"rid"`
			Msg string    `json:"msg"`
			Ts  time.Time `json:"ts"`
			U   struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"u"`
			UpdatedAt time.Time     `json:"_updatedAt"`
			Mentions  []interface{} `json:"mentions"`
			Channels  []interface{} `json:"channels"`
		} `json:"lastMessage,omitempty"`
		ResponseBy struct {
			ID            string    `json:"_id"`
			Username      string    `json:"username"`
			LastMessageTs time.Time `json:"lastMessageTs"`
		} `json:"responseBy,omitempty"`
		Metrics struct {
			Reaction struct {
				Fd time.Time `json:"fd"`
				Ft float64   `json:"ft"`
				Tt float64   `json:"tt"`
			} `json:"reaction"`
			Response struct {
				Avg   float64   `json:"avg"`
				Fd    time.Time `json:"fd"`
				Ft    float64   `json:"ft"`
				Total int       `json:"total"`
				Tt    float64   `json:"tt"`
			} `json:"response"`
		} `json:"metrics,omitempty"`
		ClosedAt time.Time `json:"closedAt,omitempty"`
		ClosedBy struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"closedBy,omitempty"`
		Closer   string `json:"closer,omitempty"`
		Metricss struct {
			ChatDuration        float64 `json:"chatDuration"`
			ServiceTimeDuration float64 `json:"serviceTimeDuration"`
			Reaction            struct {
				Fd time.Time `json:"fd"`
				Ft float64   `json:"ft"`
				Tt float64   `json:"tt"`
			} `json:"reaction"`
			Response struct {
				Avg   float64   `json:"avg"`
				Fd    time.Time `json:"fd"`
				Ft    float64   `json:"ft"`
				Total int       `json:"total"`
				Tt    float64   `json:"tt"`
			} `json:"response"`
		} `json:"metrics,omitempty"`
		Tags         []interface{} `json:"tags,omitempty"`
		LastMessagee struct {
			T                   string    `json:"t"`
			Msg                 string    `json:"msg"`
			Groupable           bool      `json:"groupable"`
			TranscriptRequested bool      `json:"transcriptRequested"`
			Ts                  time.Time `json:"ts"`
			U                   struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"u"`
			Rid       string        `json:"rid"`
			ID        string        `json:"_id"`
			UpdatedAt time.Time     `json:"_updatedAt"`
			Mentions  []interface{} `json:"mentions"`
			Channels  []interface{} `json:"channels"`
		} `json:"lastMessage,omitempty"`
		VV struct {
			ID            string    `json:"_id"`
			Username      string    `json:"username"`
			Token         string    `json:"token"`
			Status        string    `json:"status"`
			LastMessageTs time.Time `json:"lastMessageTs"`
		} `json:"v,omitempty"`
		LastMessageee struct {
			ID    string    `json:"_id"`
			Rid   string    `json:"rid"`
			Msg   string    `json:"msg"`
			Token string    `json:"token"`
			Alias string    `json:"alias"`
			Ts    time.Time `json:"ts"`
			U     struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"u"`
			UpdatedAt time.Time     `json:"_updatedAt"`
			Mentions  []interface{} `json:"mentions"`
			Channels  []interface{} `json:"channels"`
		} `json:"lastMessage,omitempty"`
		Metricsss struct {
			V struct {
				Lq time.Time `json:"lq"`
			} `json:"v"`
		} `json:"metrics,omitempty"`
	} `json:"history"`
}

type LiveChatVisitorSearchResponse struct {
	Visitors []models.LiveChatVisitor
	base
}

func (c *Client) LivechatAgentInfo(req *models.AgentInfoRequest) (*AgentInfoResponse, error) {
	response := new(AgentInfoResponse)
	if err := c.Get("livechat/agent.info"+"/"+req.RID+"/"+req.Token, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatAgentNext(req *models.AgentInfoRequest) (*AgentInfoResponse, error) {
	response := new(AgentInfoResponse)
	if err := c.Get("livechat/agent.next"+"/"+req.Token, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatAgentDepartments(params url.Values) (*AgentDepartmentResponse, error) {
	response := new(AgentDepartmentResponse)
	api := "livechat/agents/" + params["agentId"][0] + "/departments"
	if err := c.Get(api, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatAppearance() (*AppearanceResponse, error) {
	response := new(AppearanceResponse)
	if err := c.Get("livechat/appearance", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatConfig(params url.Values) (*LiveChatConfigResponse, error) {
	response := new(LiveChatConfigResponse)
	if err := c.Get("livechat/config", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatCustomField(req *LiveChatCustomFieldRequest) (*LiveChatCustomFieldResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(LiveChatCustomFieldResponse)
	err = c.Post("livechat/custom.field", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatCustomFields(req *LiveChatCustomFieldsRequest) (*LiveChatCustomFieldsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(LiveChatCustomFieldsResponse)
	err = c.Post("livechat/custom.fields", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatCustom_Fields(params url.Values) (*LiveChatCustom_FieldsResponse, error) {
	response := new(LiveChatCustom_FieldsResponse)
	if err := c.Get("livechat/custom-fields", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatCustom_Field(params url.Values) (*LiveChatCustom_FieldResponse, error) {
	response := new(LiveChatCustom_FieldResponse)
	if err := c.Get("livechat/custom-fields"+"/"+params["_id"][0], nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatDepartment() (*LiveChatDepartmentsResponse, error) {
	response := new(LiveChatDepartmentsResponse)
	if err := c.Get("livechat/department", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

//
func (c *Client) LivechatRegisterDepartment(req *LiveChatResgisterDepartmentRequest) (*LiveChatRegisterDepartmentResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(LiveChatRegisterDepartmentResponse)
	err = c.Post("livechat/department", bytes.NewBuffer(body), response)
	return response, err
}

//Get info about a department
func (c *Client) LivechatQueryDepartmentGet(params url.Values) (*LiveChatQueryDepartmentResponse, error) {
	api := "livechat/department" + "/" + params["_id"][0]
	response := new(LiveChatQueryDepartmentResponse)
	if err := c.Get(api, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

//
func (c *Client) LivechatUpdateDepartment(params url.Values, req *LiveChatUpdateDepartmentRequest) (*LiveChatUpdateDepartmentResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	api := "livechat/department" + "/" + params["_id"][0]
	response := new(LiveChatUpdateDepartmentResponse)
	err = c.Put(api, bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatDeleteDepartmentGet(params url.Values) (*LiveChatRemoveDepartmentResponse, error) {
	api := "livechat/department" + "/" + params["_id"][0]
	response := new(LiveChatRemoveDepartmentResponse)
	if err := c.Delete(api, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatListInquiries(params url.Values) (*LiveChatListInquiriesResponse, error) {
	response := new(LiveChatListInquiriesResponse)
	if err := c.Get("livechat/inquiries.list", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatTakeInquiries(req *LiveChatTakeInquiriesRequest) (*LiveChatTakeInquiriesResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(LiveChatTakeInquiriesResponse)
	err = c.Post("livechat/inquiries.take", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatGetOneInquiry(params url.Values) (*LiveChatGetOneInquiriesResponse, error) {
	response := new(LiveChatGetOneInquiriesResponse)
	if err := c.Get("livechat/inquiries.getOne", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatIntegrationSetting() (*LiveChatIntegraionSettingResponse, error) {
	response := new(LiveChatIntegraionSettingResponse)
	if err := c.Get("livechat/integrations.settings", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatMessage(req *LiveChatMessageRequest) (*LiveChatMessageResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(LiveChatMessageResponse)
	err = c.Post("livechat/message", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatUpdateMessage(params url.Values, req *LiveChatUpdateMessageRequest) (*LiveChatUpdateMessageResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	api := "livechat/message/" + params["_id"][0]
	response := new(LiveChatUpdateMessageResponse)
	err = c.Put(api, bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatRemoveMessage(params url.Values, req *LiveChatDeleteMessageRequest) (*LiveChatDeleteMessageResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(LiveChatDeleteMessageResponse)
	api := "livechat/message/" + params["_id"][0]
	if err := c.Delete(api, bytes.NewBuffer(body), response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatLoadHistory(params url.Values) (*LiveChatLoadHistoryResponse, error) {
	response := new(LiveChatLoadHistoryResponse)
	api := "livechat/messages.history/" + params["rid"][0]
	if err := c.Get(api, params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatOfflineSend(req *LiveChatOfflineSendRequest) (*LiveChatOfflineSendResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(LiveChatOfflineSendResponse)
	err = c.Post("livechat/offline.message", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatOfficeHour() (*LiveChatOfficeHourResponse, error) {
	response := new(LiveChatOfficeHourResponse)
	if err := c.Get("livechat/office-hours", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatQueue(params url.Values) (*LiveChatQueueResponse, error) {
	response := new(LiveChatQueueResponse)
	if err := c.Get("livechat/queue", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatRoom(params url.Values) (*LiveChatRoomResponse, error) {
	response := new(LiveChatRoomResponse)
	if err := c.Get("livechat/room", params, response); err != nil {
		return nil, err
	}
	return response, nil

}

func (c *Client) LivechatRoomClose(req *LiveChatRoomCloseRequest) (*LiveChatRoomCloseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(LiveChatRoomCloseResponse)
	err = c.Post("livechat/room.close", bytes.NewBuffer(body), response)
	return response, err

}

func (c *Client) LivechatRoomTransfer(req *LiveChatRoomTransferRequest) (*LiveChatRoomTransferResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(LiveChatRoomTransferResponse)
	err = c.Post("livechat/room.transfer", bytes.NewBuffer(body), response)
	return response, err

}

func (c *Client) LivechatRoomForward(req *LiveChatRoomForwardRequest) (*LiveChatRoomForwardResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(LiveChatRoomForwardResponse)
	err = c.Post("livechat/room.forward", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatRoomSurvey(req *LiveChatRoomSurveyRequest) (*LiveChatRoomSurveyResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(LiveChatRoomSurveyResponse)
	err = c.Post("livechat/room.survey", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatRooms(params url.Values) (*LiveChatRoomsResponse, error) {
	response := new(LiveChatRoomsResponse)
	if err := c.Get("livechat/rooms", params, response); err != nil {
		return nil, err
	}
	return response, nil

}

func (c *Client) LivechatTranscript(req *LiveChatTranscriptRequest) (*LiveChatTranscriptResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(LiveChatTranscriptResponse)
	err = c.Post("livechat/transcript", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatTriggers() (*LiveChatTriggersResponse, error) {
	response := new(LiveChatTriggersResponse)
	if err := c.Get("livechat/triggers", nil, response); err != nil {
		return nil, err
	}
	return response, nil

}

func (c *Client) LivechatTrigger(params url.Values) (*LiveChatTriggerResponse, error) {
	response := new(LiveChatTriggerResponse)
	api := "livechat/triggers/" + params["_id"][0]
	if err := c.Get(api, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatGetUsers(params url.Values) (*LiveChatGetUsersResponse, error) {
	response := new(LiveChatGetUsersResponse)
	api := "livechat/users/" + params["type"][0]
	if err := c.Get(api, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatRegisterUser(params url.Values, req *LiveChatRegisterUserRequset) (*LiveChatRegisterUserResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(LiveChatRegisterUserResponse)
	api := "livechat/users/" + params["type"][0]
	err = c.Post(api, bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatGetInfoUser(params url.Values) (*LiveChatGetInfoUserResponse, error) {
	response := new(LiveChatGetInfoUserResponse)
	api := "livechat/users/" + params["type"][0] + "/" + params["_id"][0]
	if err := c.Get(api, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatRemoveUser(params url.Values) (*LiveChatRemoveUserResponse, error) {
	response := new(LiveChatRemoveUserResponse)
	api := "livechat/users/" + params["type"][0] + "/" + params["_id"][0]
	if err := c.Delete(api, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatVideoCall(params url.Values) (*LiveChatVideoCallResponse, error) {
	response := new(LiveChatVideoCallResponse)
	api := "livechat/video.call/" + params["token"][0]
	newParams := url.Values{"rid": {params["rid"][0]}}
	if err := c.Get(api, newParams, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatRetrieveVisitor(params url.Values) (*LiveChatVisitorRetrieveResponse, error) {
	response := new(LiveChatVisitorRetrieveResponse)
	api := "livechat/visitor/" + params["token"][0]
	if err := c.Get(api, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatRegisterVisitor(req *LiveChatVisitorRegisterRequest) (*LiveChatVisitorRegisterResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(LiveChatVisitorRegisterResponse)
	err = c.Post("livechat/visitor", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatPageNavigationHistory(req *LiveChatVisitorNavigationHistoryRequest) (*LiveChatVisitorNavigationHistoryResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(LiveChatVisitorNavigationHistoryResponse)
	err = c.Post("livechat/page.visited", bytes.NewBuffer(body), response)
	return response, err
}

func (c *Client) LivechatVisitorInfo(params url.Values) (*LiveChatVisitorInfoResponse, error) {
	response := new(LiveChatVisitorInfoResponse)
	if err := c.Get("livechat/visitors.info", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatVisitorPageHistory(params url.Values) (*LiveChatVisitorNavigationHistoryRetrieveResponse, error) {
	response := new(LiveChatVisitorNavigationHistoryRetrieveResponse)
	if err := c.Get("livechat/visitors.pagesVisited", params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatVisitorChatHistory(params url.Values) (*LiveChatVisitorChatHistoryResponse, error) {
	response := new(LiveChatVisitorChatHistoryResponse)
	api := "livechat/visitors.chatHistory/room/" + params["roomId"][0] + "/visitor/" + params["visitorId"][0]
	if err := c.Get(api, params, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) LivechatVisitorSearch(params url.Values) (*LiveChatVisitorSearchResponse, error) {
	response := new(LiveChatVisitorSearchResponse)
	if err := c.Get("livechat/visitors.search", params, response); err != nil {
		return nil, err
	}
	return response, nil
}
