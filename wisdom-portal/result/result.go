package result

// http code 返回码说明
//  201  创建数据成功
//  412  URL传参没有满足条件
//  406  绑定数据类型失败  JSON XML FORM DATA 等
//  400  数据验证失败  Struct
//  500  服务端错误，涉及到model层操作，错误的返回
//  401  未授权登录

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
