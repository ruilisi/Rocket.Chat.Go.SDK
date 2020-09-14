package models

import "time"

type Permission struct {
	ID        string   `json:"_id"`
	UpdatedAt string   `json:"_updatedAt.$date"`
	Roles     []string `json:"roles"`
}

type Update struct {
	ID                  string        `json:"_id"`
	UpdatedAt           time.Time     `json:"_updatedAt,omitempty"`
	Group               string        `json:"group,omitempty"`
	GroupPermissionID   string        `json:"groupPermissionId,omitempty"`
	Level               string        `json:"level,omitempty"`
	Roles               []interface{} `json:"roles,omitempty"`
	Section             string        `json:"section,omitempty"`
	SectionPermissionID string        `json:"sectionPermissionId,omitempty"`
	SettingID           string        `json:"settingId,omitempty"`
	Sorter              int           `json:"sorter,omitempty"`
}
