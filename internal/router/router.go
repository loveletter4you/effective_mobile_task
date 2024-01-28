package router

import (
	"github.com/gin-gonic/gin"
	"github.com/loveletter4you/effective_mobile_task/config"
	"github.com/loveletter4you/effective_mobile_task/internal/controllers"
)

type HttpServer struct {
	router     *gin.Engine
	controller *controllers.Controller
}

func NewServer() *HttpServer {
	return &HttpServer{
		router:     gin.Default(),
		controller: controllers.NewController(),
	}
}

func (server *HttpServer) Routes() {
	server.router.Static("/static", "/static")
	api := server.router.Group("/api")
	api.POST("/person", server.controller.CreatePerson)
}

func (server *HttpServer) StartServer(cfg *config.Config) error {
	if err := server.controller.OpenConnection(cfg); err != nil {
		return err
	}
	server.Routes()
	if err := server.router.Run(":" + cfg.Server.Port); err != nil {
		return err
	}
	err := server.controller.CloseConnection()
	return err
}
