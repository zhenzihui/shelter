package model

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"projectsuika.com/shelter/src/util/file_util"
	"time"
)

var DB *gorm.DB
var sqliteDBDir = fmt.Sprintf("%s/.shelter", file_util.UserHome)
var sqliteFileName = fmt.Sprintf("%s/shelter.db", sqliteDBDir)

func ConnectDatabase() {
	os.MkdirAll(sqliteDBDir, 0700)
	if !file_util.FileExists(sqliteFileName) {
		_, err := os.Create(sqliteFileName)
		if err != nil {
			panic(err)
		}
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	database, err := gorm.Open(sqlite.Open(sqliteFileName), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database

}
