package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/heloayer/to-do-list/internal/controller/endpoint"
)

// NewRouter создает промежуточные компоненты Logger и Recovery и принимает запросы

func NewRouter(handler *gin.Engine) {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.POST("/api/todo-list/tasks", endpoint.CreateTask)

}
