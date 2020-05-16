package result

type LoginResult struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data LoginTokenResult `json:"data"`
}

type LoginTokenResult struct {
	Token string `json:"token"`
}

func NewLoginResult(code int, data LoginTokenResult) *LoginResult {
	return &LoginResult{
		Code: code,
		Msg:  ResultText(code),
		Data: data,
	}
}

func NewLoginTokenResult(token string) *LoginTokenResult {
	return &LoginTokenResult{Token: token}
}
