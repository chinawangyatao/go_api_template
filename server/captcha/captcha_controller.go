package captcha

import (
	"github.com/gin-gonic/gin"
	"go_admin/utils"
)

func ResCaptcha(c *gin.Context) {
	utils.L.Info(c)
	id, base64Image := MakeCaptcha()
	utils.Success(c, map[string]interface{}{"imageCode": id, "images": base64Image})
}
