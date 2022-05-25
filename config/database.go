package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() (string, bool) {
	dsn := "host=" + Host + " user=" + Username + " password=" + Password + " dbname=" + Database + " port=" + Port + " sslmode=disable TimeZone=Asia/Jakarta"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		return err.Error(), true
	}

	return "Postgres server listening on port " + Port + "...\n", false
}