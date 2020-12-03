package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var conn *gorm.DB
var err error

// Connection is a database connection
func Connection(){	

	// dsn TODO: comfirm at the end is "?charset=utf8mb4&parseTime=True&loc=Local" necessary
	authDB := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	if err != nil {
	logrus.Error("DB", err)
	}
	
	conn,err = gorm.Open(mysql.Open(authDB), &gorm.Config{})
	if err != nil {
		logrus.Error("DB" , err)
	}else{
		logrus.Info("Connection Opened to Database")
	}
	fmt.Printf(authDB)
}

// SeedConnection is connection for seeder
func SeedConnection(){
	Connection()
	if err := godotenv.Load(); err != nil {
		logrus.Error("Error: ", err)
	}
	logrus.Info("Connection Opened:Seeder to Database")
}

// GetDB -> method to get connection instance
func GetDB() *gorm.DB {
	return conn
}

// AppConnection -> method to get connection instance
func AppConnection() {
	if err := godotenv.Load(); err != nil {
		logrus.Error("ENV", err)
	}
	Connection()
}

// TestConnection -> method to test connection instance
func TestConnection() {
	if err := godotenv.Load("../.env"); err != nil {
		logrus.Error("ENV", err)
	}
	Connection()
}