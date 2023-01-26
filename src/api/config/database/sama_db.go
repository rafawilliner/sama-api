package database

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	maxLifetime = time.Minute * 1
	maxIdleConn = 20
	maxOpenConn = 20
)

type Connection interface {
	Connect() (client *gorm.DB, err error)
}

type GormConnection struct {
}

func (connection GormConnection) Connect() (client *gorm.DB, err error) {
	connectionData := GetConnectionDataBase()
	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	}

	client, err = gorm.Open(connectionData.DialectConnect(GetConnectionString(connectionData)), &config)

	if err != nil {
		panic(err)
	}

	sqlDB, err := client.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(maxLifetime)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)

	return client, nil
}
