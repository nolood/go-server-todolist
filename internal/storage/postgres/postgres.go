package postgres

import (
	"fmt"
	"go-server/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var Db *gorm.DB

func ConnectDb() {
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASSWORD")
	name := viper.GetString("DB_NAME")
	host := viper.GetString("DB_HOST")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		config.Logger.Info(err.Error())
	}

	err = Db.AutoMigrate(&User{}, &Article{}, &Record{}, &Bill{})
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}

	createDefaultArticles()

}
