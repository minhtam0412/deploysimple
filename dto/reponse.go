package models

type Result struct {
	err error `json:"error"`
	data interface{} `json:"data"`
}
