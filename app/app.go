package app

import (
	"fmt"
	"os"
	"simpleapi/config"
	"simpleapi/repository"
	"simpleapi/route"
	"simpleapi/service"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB)
	app := service.NewService(repo)
	route.RegisterAPI(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
