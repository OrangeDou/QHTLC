package utils

const (
	SUCCESS = 200
	ERROR   = 500

	//code = 1000... 	用户模块错误
	ERROR_USERNAME_USED              = 1001 //用户名已存在
	ERROR_USERNAME_OR_PASSWORD_WRONG = 1002 //用户名或密码错误
	ERROR_USER_NOT_FUND              = 1003 //用户名不存在
	ERROR_TOKEN_EXIT                 = 1004 //TOKEN不存在
	ERROR_TOKEN_RUNTIM               = 1005 //TOKEN已过期
	ERROR_TOKEN_WRONG                = 1006 //TOKEN不正确
	ERROR_TOKEN_TYPE_WRONG           = 1007 //TOKEN格式不正确

	//code = 2000.... 文章模块的错误

	//code = 3000 ... 其他错误
)

var CodeMsg = map[int]string{
	SUCCESS:                          "OK",
	ERROR:                            "FAIL",
	ERROR_USERNAME_USED:              "用户名已存在",
	ERROR_USERNAME_OR_PASSWORD_WRONG: "用户名或密码错误",
	ERROR_USER_NOT_FUND:              "用户名不存在",
	ERROR_TOKEN_EXIT:                 "TOKEN不存在",
	ERROR_TOKEN_RUNTIM:               "TOKEN已过期",
	ERROR_TOKEN_WRONG:                "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:           "TOKEN格式不正确",
}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}
