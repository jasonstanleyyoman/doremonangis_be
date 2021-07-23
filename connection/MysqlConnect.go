package connection

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConnection struct {
	Db gorm.DB
}

var lock = &sync.Mutex{}
var connection *MysqlConnection

func GetConnection() *MysqlConnection {
	if connection == nil {
		lock.Lock()
		defer lock.Unlock()
		if connection == nil {

			connection = &MysqlConnection{}
			connection.Connect()
		}
	}
	return connection
}

func (conn *MysqlConnection) Connect() {

	dsn := "admin:password@tcp(127.0.0.1:3306)/doremonangis?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		os.Exit(1)
	}
	conn.Db = *db
}
