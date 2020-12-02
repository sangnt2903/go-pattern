package config

import (
	"go-pattern/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func ConnectToDatabase() {
	dsn := "user=postgres password=S@ng29031998 dbname=myDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	// set db
	DBConnection = db.Debug()
	//migrate
	Migration()
}

func Migration() error {
	return DBConnection.AutoMigrate(
		models.User{},
	)
}

func GetDB() *gorm.DB {
	return DBConnection
}
