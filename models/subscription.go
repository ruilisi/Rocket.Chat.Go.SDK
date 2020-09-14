package models

type Read struct {
	Rid string `json:"rid"`
}

type UnRead struct {
	RoomId string `json:"roomId"`
}
