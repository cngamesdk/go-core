package database

import (
	"log"
	"testing"
)

func TestMysqlDb(t *testing.T) {
	dbConfig := &DbConfig{
		Host:      "your_host",
		Port:      3306,
		User:      "your_user",
		Password:  "your_password",
		Database:  "your_database",
		ParseTime: true,
		Charset:   "utf8mb4",
	}
	db, err := MysqlDb(dbConfig)

	result := map[string]interface{}{}
	db.Table("cv_platform").Take(&result)

	log.Println(db)
	log.Println(err)
	log.Println(result)

}
