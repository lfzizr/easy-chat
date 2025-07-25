package xerr

var codeText = map[int]string{
	SERVER_COMMON_ERR: "服务器内部错误",
	REQUEST_PARAM_ERR: "请求参数错误",
	DB_ERR:            "数据库错误",
}

func ErrMsg(errcode int) string {
	if msg, ok := codeText[errcode]; ok {
		return msg
	}
	return codeText[SERVER_COMMON_ERR]
}