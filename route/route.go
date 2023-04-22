package route

import (
	"simpleapi/handler"
	"simpleapi/service"

	"github.com/gin-gonic/gin"
)

func RegisterAPI(e *gin.Engine, svc service.ServiceInterface) {
	server := handler.NewHttpServer(svc)

	api := e.Group("/book")
	{
		api.POST("", server.CreateBook)
		api.GET("/:id", server.GetBookByID)
		api.GET("/all", server.GetAllBooks)
		api.PUT("/:id", server.UpdateBook)
		api.DELETE("/:id", server.DeleteBook)

	}
}
