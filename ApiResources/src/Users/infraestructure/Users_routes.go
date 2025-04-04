package infraestructure

import (
	"api_resources/src/Users/infraestructure/controllers"
	"api_resources/src/core"

	"github.com/gin-gonic/gin"
)

type UsersHandlers struct {
	create         *controllers.CreateUserController
	createAdmin    *controllers.CreateUserAdminController // New handler
	get            *controllers.GetUserByIDController
	update         *controllers.UpdateUserController
	delete         *controllers.DeleteUserController
	getAll         *controllers.GetAllUsersController
	getByEmail     *controllers.GetUserByEmailController
	login          *controllers.LoginController // 👈 nuevo campo
	getByRole      *controllers.GetUsersByRoleController
	getByCreatedBy *controllers.GetUsersByCreatedByController
}

func UsersRoutes(router *gin.Engine, handlers UsersHandlers) {
	router.POST("/users", handlers.create.Execute)
	router.POST("/users/admin", handlers.createAdmin.Execute) // New route
	router.POST("/users/login", handlers.login.Execute)

	protected := router.Group("/users")
	protected.Use(core.AuthMiddleware())

	protected.GET("/email", handlers.getByEmail.Execute) // ← ahora protegida ✅
	protected.GET("/:id", handlers.get.Execute)
	protected.GET("/role/:role", handlers.getByRole.Execute)
	protected.GET("/created-by/:created_by", handlers.getByCreatedBy.Execute)
	protected.PUT("/:id", handlers.update.Execute)
	protected.DELETE("/:id", handlers.delete.Execute)
	protected.GET("", handlers.getAll.Execute)
}
