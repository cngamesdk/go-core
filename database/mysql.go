package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDb(dbConfig *DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Charset, dbConfig.ParseTime)
	//参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDb, sqlDbErr := db.DB()
	if sqlDbErr != nil {
		return nil, sqlDbErr
	}
	if dbConfig.MaxIdleConns > 0 {
		sqlDb.SetMaxIdleConns(dbConfig.MaxIdleConns)
	}
	if dbConfig.MaxOpenConns > 0 {
		sqlDb.SetMaxOpenConns(dbConfig.MaxOpenConns)
	}
	if dbConfig.ConnMaxLifetime > 0 {
		sqlDb.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)
	}
	return db, err
}
