package models

type Roles struct {
	ID           string `json:"_id"`
	Description  string `json:"description"`
	Mandatory2Fa bool   `json:"mandatory2fa"`
	Protected    bool   `json:"protected"`
	Scope        string `json:"scope"`
}

type Role struct {
	Name        string `json:"name"`
	Scope       string `json:"scopep,omitempty" gorm:"default:'Users'"`
	Description string `json:"description,omitempty"`
}

type AddUserToRole struct {
	RoleName string `json:"roleName"`
	Username string `json:"username"`
	RoomId   string `json:"roomId,omitempty"`
}
