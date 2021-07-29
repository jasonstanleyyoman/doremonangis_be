package connection

import (
	"fmt"
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

var (
	lock       sync.Mutex = sync.Mutex{}
	connection *MysqlConnection
)

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
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
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
