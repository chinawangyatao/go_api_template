// Package db 初始化 mysql
package db

import (
	"fmt"
	"go_admin/common/config"
	"go_admin/server/entity"
	"go_admin/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GormInit() {
	config.Init()
	fmt.Println(config.Config)
	dbConfig := config.Config.Db
	dsn := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName)
	//初始化 db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("init db err:%v", err)
		return
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)
		DB = db
	}
	//// 迁移数据
	err = db.AutoMigrate(
		entity.SysUsers{},
		entity.UserCard{},
		entity.UserBanner{},
		entity.UserCase{},
		entity.UserNews{},
		entity.UserPhotoAlbum{},
		entity.UserTechnology{},
		entity.UserCardLocation{},
		entity.UserBannerList{},
		entity.UserPhotoAlbumList{},
	)
	if err != nil {
		fmt.Printf("auto migrate err:%v", err)
		utils.L.Error("auto migrate err:", err)
	}

}
