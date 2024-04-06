package userInfo

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/pkg/jwt"
	"go_admin/server/entity"
	"go_admin/utils"
)

// ISysUserInfo 定义接口
type ISysUserInfo interface {
	Login(c *gin.Context, dto entity.LoginDto)
	Register(c *gin.Context, dto entity.SysUsers)
	VerifyUsername(c *gin.Context, dto entity.VerifyRegisterDto)
}

type SysUserInfo struct {
}

// Login 登录逻辑
func (s SysUserInfo) Login(c *gin.Context, dto entity.LoginDto) {
	if dto.Username == "admin" { // 方便调试写的逻辑正式删除
		info := SysUserInfoDetails(dto) //拿到用户信息
		//	生成 token
		token, err := jwt.GenerateToken(info)
		if err != nil {
			utils.L.Error("Generate Token Err！", err)
			return
		}
		utils.Success(c, map[string]interface{}{"token": token, "userInfo": info})
		return
	}

	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		//少参数
		utils.Failed(c, int(utils.ApiCode.MISSINGLOGINPARAMS), utils.ApiCode.GetMessage(utils.ApiCode.MISSINGLOGINPARAMS))
		return
	}

	//	验证码是否过期
	code := utils.RedisStore{}.Get(dto.ImageCode, true)
	if len(code) == 0 {
		utils.Failed(c, int(utils.ApiCode.CAPTCODEEXPIRED), utils.ApiCode.GetMessage(utils.ApiCode.CAPTCODEEXPIRED))
		return

	}

	//	校验验证码
	verifyRes := utils.RedisStore{}.Verify(dto.ImageCode, dto.Captcha, true)
	if !verifyRes {
		utils.Failed(c, int(utils.ApiCode.CAPYCHANOTTRUE), utils.ApiCode.GetMessage(utils.ApiCode.CAPYCHANOTTRUE))
		return
	}

	//	校验
	info := SysUserInfoDetails(dto) //拿到用户信息
	if info.Password != dto.Password {
		utils.Failed(c, 405, "账号或者密码错误")
		return
	}

	//	生成 token
	token, err := jwt.GenerateToken(info)
	if err != nil {
		utils.L.Error("Generate Token Err！", err)
		return
	}
	utils.Success(c, map[string]interface{}{"token": token, "userInfo": info})

}

// Register 注册逻辑
func (s SysUserInfo) Register(c *gin.Context, dto entity.SysUsers) {
	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		//少参数
		utils.Failed(c, int(utils.ApiCode.MISSINGLOGINPARAMS), utils.ApiCode.GetMessage(utils.ApiCode.MISSINGLOGINPARAMS))
		return
	}

	//	将信息存入数据库
	info, ok := SysUserInfoSave(dto)
	if ok != nil {
		utils.Failed(c, 201, ok.Error())
		utils.L.Error(ok)
		return
	}
	utils.L.Info("注册成功")

	//	生成 token
	token, err := jwt.GenerateToken(info)
	if err != nil {
		utils.L.Error("Generate Token Err！", err)
		return
	}
	utils.Success(c, map[string]interface{}{"token": token, "userInfo": info})
}

// VerifyUsername 校验用户名
func (s SysUserInfo) VerifyUsername(c *gin.Context, userinfo entity.VerifyRegisterDto) {
	info := VerifyInfoUsername(userinfo)
	if info.Username != "" {
		utils.Failed(c, 201, "用户名已经存在，请换一个！")
		return
	}

	if info.Email != "" {
		utils.Failed(c, 201, "邮箱已经被占用，请换一个！")
		return
	}
	utils.Failed(c, 200, "验证通过！")
}

var sysUserInfoService = SysUserInfo{}

func SysUserInfoServiceFunc() ISysUserInfo {
	return &sysUserInfoService
}
