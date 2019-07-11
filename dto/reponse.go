package dto

type Result struct {
	Data         interface{} `json:"result"`
	ErrorMessage string      `json:"error-message"`
}
