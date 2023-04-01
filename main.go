package main

import (
	"simpleapi/app"
	"simpleapi/config"
	//"simpleapi/routers"

	"github.com/joho/godotenv"
)

func init() {

	godotenv.Load()

	err := config.InitProgress()
	if err != nil {
		panic(err)
	}

}

func main() {
	app.StartApplication()
}
