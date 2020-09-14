package models

type Setting struct {
	ID           string  `json:"_id"`
	Blocked      bool    `json:"blocked"`
	Group        string  `json:"group"`
	Hidden       bool    `json:"hidden"`
	Public       bool    `json:"public"`
	Type         string  `json:"type"`
	PackageValue string  `json:"packageValue"`
	Sorter       int     `json:"sorter"`
	Value        string  `json:"value"`
	ValueBool    bool    `json:"valueBool"`
	ValueInt     float64 `json:"valueInt"`
	ValueSource  string  `json:"valueSource"`
	ValueAsset   Asset   `json:"asset"`
}

type Asset struct {
	DefaultUrl string `json:"defaultUrl"`
}

type SettingsMessage struct {
	ID    string `json:"_id"`
	Value bool   `json:"value"`
}

type SettingsUpdate struct {
	ID      string `json:"_id"`
	Value   bool   `json:"value"`
	Color   string `json:"color"`
	Editor  bool   `json:"editor"`
	Execute bool   `json:"execute"`
}

type SettingsOauth struct {
	ID                   string `json:"_id"`
	Name                 string `json:"name"`
	Service              string `json:"service"`
	ClientID             string `json:"clientId"`
	Custom               bool   `json:"custom"`
	ServerURL            string `json:"serverURL"`
	TokenPath            string `json:"tokenPath"`
	IdentityPath         string `json:"identityPath"`
	AuthorizePath        string `json:"authorizePath"`
	Scope                string `json:"scope"`
	ButtonLabelText      string `json:"buttonLabelText"`
	ButtonLabelColor     string `json:"buttonLabelColor"`
	LoginStyle           string `json:"loginStyle"`
	ButtonColor          string `json:"buttonColor"`
	TokenSentVia         string `json:"tokenSentVia"`
	IdentityTokenSentVia string `json:"identityTokenSentVia"`
	UsernameField        string `json:"usernameField"`
	MergeUsers           bool   `json:"mergeUsers"`
}

type SettingsPrivate struct {
	ID    string `json:"_id"`
	Value bool   `json:"value"`
}
