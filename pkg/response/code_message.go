package response

type ResCode int64

const (
	CodeSuccess ResCode = 2000
)

const (
	CodeUserExist ResCode = 3000 + iota
	CodeUserNotExist
	CodeInvalidToken

	CodeNeedLogin
	CodeInvalidPassword
	CodeServerBusy
	CodeInvalidParam
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
	CodeUserExist: "用户名已存在",
	CodeUserNotExist: "用户名不存在",
	CodeInvalidToken: "无效的token",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:"系统繁忙",
	CodeNeedLogin: "需要登录",
	CodeInvalidParam: "请求参数错误",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

