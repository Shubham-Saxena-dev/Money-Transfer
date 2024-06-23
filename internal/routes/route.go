package routes

/*
This class is a route handler. From here, the requests are directed towards the controller
*/

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qonto/pkg/handlers"
)

type Route interface {
	RegisterHandlers()
}

type route struct {
	engine     *gin.Engine
	controller handlers.AccountController
}

func RegisterHandlers(engine *gin.Engine, controller handlers.AccountController) Route {
	return &route{
		engine:     engine,
		controller: controller,
	}
}

// RegisterHandlers This is a route handler for various requests
func (r *route) RegisterHandlers() {
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.engine.POST("/transfer", r.controller.TransferHandler)
}
