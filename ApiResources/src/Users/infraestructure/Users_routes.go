package infraestructure

import (
    "api_resources/src/Users/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

type UsersHandlers struct {
    create *controllers.CreateUserController
    get   	*controllers.GetUserByIDController
    update *controllers.UpdateUserController
    delete *controllers.DeleteUserController
    getAll *controllers.GetAllUsersController
	getByEmail *controllers.GetUserByEmailController
}

func UsersRoutes(router *gin.Engine, handlers UsersHandlers) {
    router.POST("/users", handlers.create.Execute)
    router.GET("/users/:id", handlers.get.Execute)
    router.PUT("/users/:id", handlers.update.Execute)
    router.DELETE("/users/:id", handlers.delete.Execute)
    router.GET("/users", handlers.getAll.Execute)
	router.GET("/users/email", handlers.getByEmail.Execute)
}