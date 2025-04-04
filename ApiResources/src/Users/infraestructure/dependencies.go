package infraestructure

import (
	"api_resources/src/Users/application"
	"api_resources/src/Users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	os := NewPostgreSQL()

	// Casos de uso
	createUserService := application.NewCreateUserUseCase(os)
	createUserAdminService := application.NewCreateUserAdminUseCase(os) // New use case
	getUserService := application.NewGetUserByID(os)
	updateUserService := application.NewUpdateUser(os)
	deleteUserService := application.NewDeleteUser(os)
	getAllUsersService := application.NewGetAllUsers(os)
	getUserByEmailService := application.NewGetUserByEmail(os)
	loginService := application.NewLoginUseCase(os)
	getUsersByRoleService := application.NewGetUsersByRole(os)
	getUsersByCreatedByService := application.NewGetUsersByCreatedByUseCase(os)

	// Controladores
	createUserController := controllers.NewCreateUserController(*createUserService)
	createUserAdminController := controllers.NewCreateUserAdminController(createUserAdminService) // New controller
	getUserController := controllers.NewGetUserByIDController(*getUserService)
	updateUserController := controllers.NewUpdateUserController(*updateUserService)
	deleteUserController := controllers.NewDeleteUserController(*deleteUserService)
	getAllUsersController := controllers.NewGetAllUsersController(*getAllUsersService)
	getUserByEmailController := controllers.NewGetUserByEmailController(*getUserByEmailService)
	loginController := controllers.NewLoginController(loginService)
	getUsersByRoleController := controllers.NewGetUsersByRoleController(*getUsersByRoleService)
	getUsersByCreatedByController := controllers.NewGetUsersByCreatedByController(getUsersByCreatedByService)

	UsersRoutes(router, UsersHandlers{
		create:         createUserController,
		createAdmin:    createUserAdminController, // New handler
		get:            getUserController,
		update:         updateUserController,
		delete:         deleteUserController,
		getAll:         getAllUsersController,
		getByEmail:     getUserByEmailController,
		login:          loginController,
		getByRole:      getUsersByRoleController,
		getByCreatedBy: getUsersByCreatedByController,
	})
}
