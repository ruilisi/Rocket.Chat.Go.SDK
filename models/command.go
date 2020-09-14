package models

type Command struct {
	Command     string `json:"command"`
	Params      string `json:"params"`
	Description string `json:"description"`
	ClientOnly  bool   `json:"clientOnly"`
}
