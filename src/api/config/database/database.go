package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mysqlDialect   = "mysql"
	sqlLiteDialect = "sqlite3"
)

type ConnectionData struct {
	Host           string
	Schema         string
	Username       string
	Password       string
	Dialect        string
	DialectConnect func(dns string) gorm.Dialector
}

func GetConnectionDataBase() *ConnectionData {

	connectionData := ConnectionData{}
	connectionData.Host = os.Getenv("DB_HOST")
	connectionData.Schema = os.Getenv("DB_SCHEMA")
	connectionData.Username = os.Getenv("DB_USER")
	connectionData.Password = os.Getenv("DB_PASS")
	connectionData.Dialect = mysqlDialect
	connectionData.DialectConnect = mySQLConnect
	return &connectionData
}

func mySQLConnect(dns string) gorm.Dialector {
	return mysql.Open(dns)
}

func GetConnectionString(cd *ConnectionData) string {
	if cd.Dialect == sqlLiteDialect {
		return cd.Host
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", cd.Username, cd.Password, cd.Host, cd.Schema)
}
