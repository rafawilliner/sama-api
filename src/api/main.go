package main

import (
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/rafawilliner/sama-api/src/api/app"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load(filepath.Join("./src/api/", ".env"))
		if err != nil {
			panic(err)
		}
	}

	app.Start()
}
