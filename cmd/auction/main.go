package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/natansa/go-leilao/configuration/database/mongodb"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return
	}

	ctx := context.Background()
	_, err := mongodb.NewMongoDBConnection(ctx)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
