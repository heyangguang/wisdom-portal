package result

type SuccessResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type FailResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  string `json:"err"`
}

type SliceFailResult struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Err  []map[string]string `json:"err"`
}

func NewSuccessResult(code int) *SuccessResult {
	return &SuccessResult{
		Code: code,
		Msg:  ResultText(code),
	}
}

func NewFailResult(code int, err string) *FailResult {
	return &FailResult{
		Code: code,
		Msg:  ResultText(code),
		Err:  err,
	}
}

func NewSliceFailResult(code int, err []map[string]string) *SliceFailResult {
	return &SliceFailResult{
		Code: code,
		Msg:  ResultText(code),
		Err:  err,
	}
}
