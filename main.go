package main

import (
	"go_admin/pkg/db"
	"go_admin/pkg/redis"
	router2 "go_admin/router"
	log "go_admin/utils"
)

func main() {
	log.InitLogger() // 初始化 log
	router := router2.InitRouter()

	//	启动服务
	router.Run(":9091")

}

// 初始化链接
func init() {
	//	mysql
	db.GormInit()
	//	redis
	err := redis.RedisInit()
	if err != nil {
		log.L.Error("init redis err: %s", err)
		return
	}
}
