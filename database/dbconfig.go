package database

import "time"

type DbConfig struct {
	User            string        `json:"user"`
	Password        string        `json:"password"`
	Host            string        `json:"host"`
	Port            uint32        `json:"port"`
	Database        string        `json:"database"`
	Charset         string        `json:"charset"`
	ParseTime       bool          `json:"parse_time"`
	MaxIdleConns    int           `json:"max_idle_conns"`    //MaxIdleConns 设置空闲连接池中连接的最大数量
	MaxOpenConns    int           `json:"max_open_conns"`    //MaxOpenConns 设置打开数据库连接的最大数量
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"` //ConnMaxLifetime 设置了连接可复用的最大时间
}
