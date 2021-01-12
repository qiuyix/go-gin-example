package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-gin-example/pkg/setting"
	"log"
)

var db *gorm.DB

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePerfix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal("failed to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePerfix = sec.Key("TABLE_PREFIX").String()

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Fatal(err)
	}

	// 为表明添加前缀
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePerfix + defaultTableName;
	}

	// 禁用复数形式
	db.SingularTable(true)
	// 设置打印日志
	db.LogMode(true)
	// 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}