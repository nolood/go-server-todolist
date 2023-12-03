package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConnectDb() *pg.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatal("Can't get DB_PORT")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("Can't get DB_USER")
	}
	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		log.Fatal("Can't get DB_PASSWORD")
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		log.Fatal("Can't get DB_NAME")
	}

	db := pg.Connect(&pg.Options{
		Addr:     port,
		User:     user,
		Password: pass,
		Database: name,
	})

	return db
}

func CreateSchemas() {
	models := []interface{}{
		(*User)(nil),
	}

	log.Println("Creating schemas...")

	for _, model := range models {
		err := ConnectDb().Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		log.Println("Successfully created table: ", model)
		if err != nil {
			log.Fatal(err)
		}
	}
}
