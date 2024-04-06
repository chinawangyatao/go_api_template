package captcha

import (
	"github.com/mojocn/base64Captcha"
	"go_admin/utils"
)

var store = utils.RedisStore{}

// CaptchaMake 生成验证码
func MakeCaptcha() (string, string) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	//	配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captacha := base64Captcha.NewCaptcha(driver, store)
	lid, lib64s, _, _ := captacha.Generate()
	return lid, lib64s
}
