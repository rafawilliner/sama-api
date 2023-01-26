package database

import (
	"fmt"

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
	connectionData.Host = "127.0.0.1"    //os.Getenv("DB_TEST_HOST")
	connectionData.Schema = "sama"       //os.Getenv("DB_TEST_SCHEMA")
	connectionData.Username = "root"     //os.Getenv("DB_TEST_USER")
	connectionData.Password = "mamincho" //os.Getenv("DB_TEST_PASS")
	connectionData.Dialect = mysqlDialect
	connectionData.DialectConnect = mySQLConnect
	return &connectionData
}

/*func (cd *ConnectionData) setupMasterConnectionData() *ConnectionData {
	cd.Host = "DB_MYSQL_CREDITSCREDIT00_CREDPRODS_CREDPRODS_ENDPOINT"
	cd.Password = "DB_MYSQL_CREDITSCREDIT00_CREDPRODS_CREDPRODS_WPROD"
	cd.Username = "credprods_WPROD"
	cd.Schema = "credprods"
	cd.Dialect = mysqlDialect
	cd.DialectConnect = mySQLConnect
	return cd
}*/

func mySQLConnect(dns string) gorm.Dialector {
	return mysql.Open(dns)
}

func GetConnectionString(cd *ConnectionData) string {
	if cd.Dialect == sqlLiteDialect {
		return cd.Host
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", cd.Username, cd.Password, cd.Host, cd.Schema)
}
