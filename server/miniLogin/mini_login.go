package miniLogin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"go_admin/common"
	"go_admin/server/entity"
	"go_admin/utils"
	"net/http"
	"time"
)

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 24 * 7 * 365

func Login(c *gin.Context) {
	LoginData := struct { // 定义绑定的结构体
		Code string `json:"code"`
	}{}

	err := c.BindJSON(&LoginData)
	if err != nil {
		utils.L.Error("bindjson loginData err:", err)
		return
	}
	//请求微信 url
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=%s",
		common.Appid, common.VxSecret, LoginData.Code, common.GrantType,
	)

	response, err := http.Get(url)
	if err != nil {
		utils.L.Error("Get weixin url err:", err)
		return
	}
	//获取 openid 如果需要就存下
	wxRes := WXLoginRes{}
	if err := json.NewDecoder(response.Body).Decode(&wxRes); err != nil {
		utils.L.Error("json wxRes err:", err)
		utils.Failed(c, 500, wxRes.Errmsg)
		return
	}
	info := entity.SysUsers{
		OpenId: wxRes.Openid,
	}
	miniInfo := MiniUserInfo(info) // 存下 openid

	//根据 openid 等 userInfo 生成数据
	t, err := GenerateToken(info)
	if err != nil {
		utils.L.Error("GenerateToken err:", err)
		return
	}
	utils.Success(c, map[string]interface{}{"userinfo": miniInfo, "token": t})

}

// GenerateToken 根据用户信息生成 token
func GenerateToken(jwtInfo entity.SysUsers) (string, error) {
	c := struct {
		entity.SysUsers
		jwt.RegisteredClaims
	}{
		SysUsers: jwtInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //过期时间
			Issuer:    "admin",                                                 //签发人
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signed, err := token.SignedString(common.Secret)
	return signed, err
}
