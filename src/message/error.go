package message

var ERROR_MESSAGE interface{}


func init()  {
	ERROR_MESSAGE = map[string]interface{} {
		"10000" : "成功",
		"10001" : "注册失败",
		"10002" : "登陆失败",
		"10003" : "",
	}
}