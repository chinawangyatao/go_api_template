package common

// Secret token 密钥
var Secret = []byte("go-api")
var (
	ErrAbsent  = "token absent 令牌不存在" //令牌不存在
	ErrInvalid = "token invalid 令牌无效" // 令牌无效
)

// 小程序 url
var (
	WeichatURL = "https://api.weixin.qq.com/sns/jscode2session"
	Appid      = "wxe9e884e243401dfb"
	VxSecret   = "3277a4dd79059fa33e33e36a106ee49b"
	GrantType  = "authorization_code"
)
