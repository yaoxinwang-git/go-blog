package models

import "time"

type User struct {
	Uid      int       `json:"uid,omitempty"`
	UserName string    `json:"userName,omitempty"`
	Passwd   string    `json:"passwd,omitempty"`
	Avatar   string    `json:"avatar,omitempty"` //头像
	CreateAt time.Time `json:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}


type UserInfo struct{
	Uid      int       `json:"uid,omitempty"`
	UserName string    `json:"userName,omitempty"`
	Avatar   string    `json:"avatar,omitempty"` //头像
}