package router

import (
	"githhub.com/Kr-Harshit/golang-react-todo/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	config := cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}
	router.Use(cors.New(config))

	// adding handler
	router.GET("/api/todo", middleware.GetAllTasks)
	router.POST("/api/todo", middleware.CreateTask)
	router.PUT("api/todo/update/:id", middleware.UpdateTask)
	// router.PUT("api/todo/toggle/:id", middleware.ToggleTask)
	router.DELETE("api/todo/:id", middleware.DeleteTask)
	router.DELETE("api/todo", middleware.DeleteAllTasks)

	return router
}
