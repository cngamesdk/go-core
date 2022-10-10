package database

import "time"

type DbConfig struct {
	User            string
	Password        string
	Host            string
	Port            uint32
	Database        string
	Charset         string
	ParseTime       bool
	MaxIdleConns    int           //MaxIdleConns 设置空闲连接池中连接的最大数量
	MaxOpenConns    int           //MaxOpenConns 设置打开数据库连接的最大数量
	ConnMaxLifetime time.Duration //ConnMaxLifetime 设置了连接可复用的最大时间
}
