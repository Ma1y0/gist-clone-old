package main

import (
	"github.com/Ma1y0/gist-clone/model"
	"github.com/Ma1y0/gist-clone/router"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	model.EstablishDBConnection()

	router := router.ConstructRouter()

	router.Run(":8080")
}
