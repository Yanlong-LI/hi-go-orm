package orm

import (
	"database/sql"
	logger "github.com/yanlong-li/hi-go-logger"
)

var db *sql.DB

//ConfigDb 配置数据库
func ConfigDb(driverName, dsn string) {
	var err error
	db, err = sql.Open(driverName, dsn)
	if err != nil {
		logger.Warning("sql open fail.", 0, err)
	}
}

func GetDb() *sql.DB {
	return db
}
