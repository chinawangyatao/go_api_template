// Package utils 统一状态码
package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Codes struct {
	SUCCESS            uint
	FAILED             uint
	Message            map[uint]string
	NOAUTH             uint
	AUTHFORMATERROR    uint
	MISSINGLOGINPARAMS uint
	CAPTCODEEXPIRED    uint
	CAPYCHANOTTRUE     uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:            200,
	FAILED:             501,
	NOAUTH:             401,
	AUTHFORMATERROR:    405,
	MISSINGLOGINPARAMS: 405,
	CAPTCODEEXPIRED:    405,
	CAPYCHANOTTRUE:     405,
}

// 状态信息

func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:            "成功",
		ApiCode.FAILED:             "失败",
		ApiCode.NOAUTH:             "请求头中 token 为空",
		ApiCode.AUTHFORMATERROR:    "请求头中 token 格式错误",
		ApiCode.MISSINGLOGINPARAMS: "缺少登录参数",
		ApiCode.CAPTCODEEXPIRED:    "验证码过期",
		ApiCode.CAPYCHANOTTRUE:     "验证码不正确",
	}
}

// GetMessage 出口函数
func (c Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}

// Result Package middleware 通用返回结构
type Result struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` //提示信息
	Data    interface{} `json:"data"`    // 返回信息
}

type Response struct{}

// Success 返回成功
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Message = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data
	c.JSON(http.StatusOK, res)

}

// Failed 返回失败
func Failed(c *gin.Context, code int, message string) {
	res := Result{}
	res.Code = code
	res.Message = message
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}
