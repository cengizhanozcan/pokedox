package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"pokemon/model/domain"
)

type Database struct {
	Instance *gorm.DB
}

func NewInMemoryDb() *Database {
	return &Database{Instance: initDb()}
}

func initDb() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=123 dbname=pokedex host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Exception occurred when db connection established!", err)
	}

	db.AutoMigrate(&domain.PokemonType{})
	db.AutoMigrate(&domain.PokemonImage{})
	db.AutoMigrate(&domain.PokemonBase{})
	db.AutoMigrate(&domain.PokemonProfile{})
	db.AutoMigrate(&domain.Pokemon{})
	db.AutoMigrate(&domain.Trainer{})

	return db
}
