package models

type LoginRes struct {
	Token    string   `json:"token,omitempty"`
	UserInfo UserInfo `json:"userInfo,omitempty"`
}
