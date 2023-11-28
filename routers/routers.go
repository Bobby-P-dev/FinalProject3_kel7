package routers

import (
	"github.com/Bobby-P-dev/FinalProject3_kel7/controllers"
	"github.com/Bobby-P-dev/FinalProject3_kel7/middleware"
	"github.com/gin-gonic/gin"
)

func StarApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.CreatAcount)
		userRouter.POST("/login", controllers.Login)
		userRouter.PUT("/update-account", middleware.Authentication(), controllers.PutAcount)
		userRouter.DELETE("/delete-account", middleware.Authentication(), controllers.DelletAcount)
	}
	categoryRouter := r.Group("/category")
	{
		categoryRouter.Use(middleware.Authentication())
		categoryRouter.POST("/post", middleware.RoleAuthorization(), controllers.CreateCategories)
		categoryRouter.GET("/get", controllers.GetteCategories)
		categoryRouter.PATCH("/patch/:categoryId", middleware.RoleAuthorization(), controllers.PatchCategories)
		categoryRouter.DELETE("/delete/:categoryId", middleware.RoleAuthorization(), controllers.DeleteCategories)
	}
	tasksRouter := r.Group("/tasks")
	{
		tasksRouter.Use(middleware.Authentication())
		tasksRouter.POST("/post", controllers.CreateTask)
		tasksRouter.GET("/get", controllers.GetTask)
		tasksRouter.PUT("/put/:taskId", middleware.TaskAuth(), controllers.PutTask)
		tasksRouter.PATCH("/patch-status/:taskId", middleware.TaskAuth(), controllers.PatchStatusTask)
		tasksRouter.PATCH("/patch-category/:taskId", middleware.TaskAuth(), controllers.PatchCategoryTask)
		tasksRouter.DELETE("/delete/:taskId", middleware.TaskAuth(), controllers.DeleteteTask)
	}
	return r
}
