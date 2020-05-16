package result

import "wisdom-portal/models"

type PubCurrentUserResult struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg"`
	Data models.PubCurrentUser `json:"data"`
}

type RegisterUserResult struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data models.GoogleAuth `json:"data"`
}

func NewPubCurrentUserResult(code int, data models.PubCurrentUser) *PubCurrentUserResult {
	return &PubCurrentUserResult{
		Code: code,
		Msg:  ResultText(code),
		Data: data,
	}
}

func NewRegisterUserResult(code int, data models.GoogleAuth) *RegisterUserResult {
	return &RegisterUserResult{
		Code: code,
		Msg:  ResultText(code),
		Data: data,
	}
}
