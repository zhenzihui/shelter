package response

type CodeDesc struct {
	Code string
	Desc string
}

var CodeRegistry = struct {
	OK              CodeDesc
	NoAuth          CodeDesc
	InvalidToken    CodeDesc
	InvalidPassword CodeDesc
}{
	OK:              CodeDesc{Code: "0", Desc: "成功"},
	NoAuth:          CodeDesc{Code: "4000", Desc: "没有登录"},
	InvalidToken:    CodeDesc{Code: "4001", Desc: "令牌失效"},
	InvalidPassword: CodeDesc{Code: "4002", Desc: "密码错误"},
}

type Resp struct {
	Code string `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// Data 返回数据
func Data(msg string, data any) Resp {
	return Resp{
		Code: CodeRegistry.OK.Code,
		Msg:  msg,
		Data: data,
	}
}

// Ok 返回消息
func Ok(msg string) Resp {
	return Resp{
		Code: CodeRegistry.OK.Code,
		Msg:  msg,
		Data: nil,
	}
}

func Error(code CodeDesc, msg string) Resp {
	message := code.Desc
	if len(msg) > 0 {
		message = msg
	}
	return Resp{
		Code: code.Code,
		Data: nil,
		Msg:  message,
	}
}
