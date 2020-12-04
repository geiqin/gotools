package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type DbConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Prefix   string `json:"prefix"`
}

func CreateMysqlDB(cfg *DbConfig) *gorm.DB {
	serverAddr := cfg.Host + ":" + cfg.Port

	//当前有效两种
	//connString := cfg.Username + ":" + cfg.Password + "@tcp(" + serverAddr + ")/" + cfg.Database + "?charset=utf8mb4&loc=Local"
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local", cfg.Username, cfg.Password, serverAddr, cfg.Database)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: cfg.Prefix,   // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		log.Println("mysql database connection failed :", cfg.Database)
		return nil
	}

	return db
}
