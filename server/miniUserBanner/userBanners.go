package miniUserBanner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/server/parseJwt"
	"go_admin/utils"
)

func GetBanner(c *gin.Context) {
	// 解析用户信息
	// 新建结构体
	// 查数据绑定结构体
	// 发回去
	userInfo, err := parseJwt.ParseMiniJwt(c)
	if err != nil {
		fmt.Println(err)
		utils.L.Error("GetBanner parseJwt err:", err.Error())
		return
	}

	//err := c.ShouldBind(&bannerList)
	//if err != nil {
	//	fmt.Printf("GetBanner Bind err: %v", err)
	//	utils.L.Error("GetBanner Bind err:", err)
	//	return
	//}

}

func SetUpDateBanner(c *gin.Context) {

}

func DeleteBanner(c *gin.Context) {

}
