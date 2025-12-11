package sql

import "gorm.io/gorm"

func GetFindSql(db *gorm.DB) string {
	return db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		var result []interface{}
		return tx.Find(&result)
	})
}

func GetTakeSql(db *gorm.DB) string {
	return db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		var result map[string]interface{}
		return tx.Take(&result)
	})
}
