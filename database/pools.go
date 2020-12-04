package database

import (
	"gorm.io/gorm"
	"log"
)

var pools map[string]*gorm.DB
var poolIndex []string

//数据库访问池子
func DbPools(cfg *DbConfig, max ...int) *gorm.DB {
	maxLen := 50
	if max != nil {
		maxLen = max[0]
	}
	if pools == nil {
		pools = make(map[string]*gorm.DB)
		poolIndex = make([]string, 0)
	}
	db, ok := pools[cfg.Database]
	if !ok {
		db = CreateMysqlDB(cfg)
		if db != nil {
			pools[cfg.Database] = db
			poolIndex = append(poolIndex, cfg.Database)
		}
	}
	if len(poolIndex) > maxLen {
		log.Println("db pool is fulled :", len(poolIndex))
	}
	if db == nil || db.Error != nil {
		log.Println("database is closed :", cfg.Database)
		log.Println("database is error :", db.Error.Error())
	}

	return db
}
