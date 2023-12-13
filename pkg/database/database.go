// Package database 数据库操作
package database

import (
	"database/sql"
	"errors"
	"github.com/xian137/layout-go/config"
	"github.com/xian137/layout-go/pkg/console"
	"github.com/xian137/layout-go/pkg/logger"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// DB 对象
var DB *gorm.DB
var SqlDB *sql.DB

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, _logger gormLogger.Interface) {
	// 使用 gorm.Open 连接数据库
	var err error
	gormConfig := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	}
	if _logger.(logger.GormLogger).ZapLogger != nil {
		gormConfig.Logger = _logger
	}
	DB, err = gorm.Open(dbConfig, gormConfig)
	if err != nil {
		console.ExitIf(err)
	}

	// 获取底层的 sqlDB
	SqlDB, err = DB.DB()
	if err != nil {
		console.ExitIf(err)
	}
}

func Shutdown() {
	err := SqlDB.Close()
	console.ExitIf(err)
}

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	err := stmt.Parse(obj)
	if err != nil {
		console.ExitIf(err)
	}
	return stmt.Schema.Table
}

func DeleteAllTables() error {
	var err error
	switch config.Get().Database.Connection {
	case "mysql":
		err = deleteMySQLTables()
	case "postgres":
		err = deletePostgresTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteMySQLTables() error {
	var tables []string
	dbname := CurrentDatabase()

	// 读取所有数据表
	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	DB.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		err = DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	// 开启 MySQL 外键检测
	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}

func deletePostgresTables() error {
	return DB.Exec("drop schema public cascade;create schema public;").Error
}
