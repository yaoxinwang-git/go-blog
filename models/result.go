package models

type Result struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Code  int         `json:"code,omitempty"`

}
