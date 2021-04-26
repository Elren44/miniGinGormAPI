package router

import (
	"github.com/Evelon44/MyTests/GormAPI/handlers"
	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	var router = gin.Default()

	ginGroup := router.Group("/ginGroup/")
	{
		ginGroup.GET("list", handlers.GetAllToDo)
		ginGroup.POST("add", handlers.PostNewToDo)
		ginGroup.GET("list/:id", handlers.GetToDoById)
		ginGroup.PUT("update/:id", handlers.UpdateToDoById)
		ginGroup.DELETE("remove/:id", handlers.DeleteToDoById)

	}
	return router
}
