package postgres

import (
	"go-server/internal/config"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/spf13/viper"
)

var Db *pg.DB

func ConnectDb() {
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASSWORD")
	name := viper.GetString("DB_NAME")

	Db = pg.Connect(&pg.Options{
		Addr:     ":" + port,
		User:     user,
		Password: pass,
		Database: name,
	})
}

func CreateSchemas() {
	models := []interface{}{
		(*User)(nil),
	}

	config.Logger.Info("Creating tables...")

	for _, model := range models {
		err := Db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			config.Logger.Error(err.Error())
		}
	}
}
