package postgres

import (
	"fmt"
	"go-server/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"

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

	err = Db.AutoMigrate(&User{}, &Article{}, &Record{}, &Bill{}, &RecordType{})
	if err != nil {
		config.Logger.Error(err.Error())
		return
	}

	createDefaultArticles()
	createDefaultRecordTypes()
}

type Model struct {
	ID        uint64         `gorm:"primarykey,autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
