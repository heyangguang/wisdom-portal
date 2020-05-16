package result

type CreateQrCodeResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func NewCreateQrCodeResult(code int, data string) *CreateQrCodeResult {
	return &CreateQrCodeResult{
		Code: code,
		Msg:  ResultText(code),
		Data: data,
	}
}
