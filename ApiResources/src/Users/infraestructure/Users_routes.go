package infraestructure

import (
	"api_resources/src/Users/infraestructure/controllers"
	"api_resources/src/core"

	"github.com/gin-gonic/gin"
)

type UsersHandlers struct {
	create     *controllers.CreateUserController
	get        *controllers.GetUserByIDController
	update     *controllers.UpdateUserController
	delete     *controllers.DeleteUserController
	getAll     *controllers.GetAllUsersController
	getByEmail *controllers.GetUserByEmailController
	login      *controllers.LoginController // 👈 nuevo campo
	getByRole  *controllers.GetUsersByRoleController
}

func UsersRoutes(router *gin.Engine, handlers UsersHandlers) {
	router.POST("/users", handlers.create.Execute)
	router.POST("/users/login", handlers.login.Execute)

	protected := router.Group("/users")
	protected.Use(core.AuthMiddleware())

	protected.GET("/email", handlers.getByEmail.Execute) // ← ahora protegida ✅
	protected.GET("/:id", handlers.get.Execute)
	protected.GET("/role/:role", handlers.getByRole.Execute)
	protected.PUT("/:id", handlers.update.Execute)
	protected.DELETE("/:id", handlers.delete.Execute)
	protected.GET("", handlers.getAll.Execute)
}
