package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
}

// OpenMysql 打开Mysql链接
func OpenMysql(config MySql) (resp *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.Username,
		config.Password,
		config.Path,
		config.Port,
		config.DbName,
		config.Config,
	)
	db, openErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if openErr != nil {
		err = openErr
		return
	}
	sqlDb, sqlDbErr := db.DB()
	if sqlDbErr != nil {
		err = sqlDbErr
		return
	}
	sqlDb.SetMaxOpenConns(config.MaxOpenConns)
	sqlDb.SetMaxIdleConns(config.MaxIdleConns)
	resp = db
	return
}
