package infrastructure

import (
    "api_resources/src/NfcCardAssignments/application"
    "api_resources/src/NfcCardAssignments/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
    os := NewPostgreSQL()

    // Use cases
    createAssignmentService := application.NewCreateAssignment(os)
    getAssignmentByIDService := application.NewGetAssignmentByID(os)
    getAssignmentsByUserIDService := application.NewGetAssignmentsByUserID(os)
    getAssignmentByCardUIDService := application.NewGetAssignmentByCardUID(os)
    updateAssignmentService := application.NewUpdateAssignment(os)
    deactivateAssignmentService := application.NewDeactivateAssignment(os)
    getAllAssignmentsService := application.NewGetAllAssignments(os)

    // Controllers
    createAssignmentController := controllers.NewCreateAssignmentController(createAssignmentService)
    getAssignmentByIDController := controllers.NewGetAssignmentByIDController(getAssignmentByIDService)
    getAssignmentsByUserIDController := controllers.NewGetAssignmentsByUserIDController(getAssignmentsByUserIDService)
    getAssignmentByCardUIDController := controllers.NewGetAssignmentByCardUIDController(getAssignmentByCardUIDService)
    updateAssignmentController := controllers.NewUpdateAssignmentController(updateAssignmentService)
    deactivateAssignmentController := controllers.NewDeactivateAssignmentController(deactivateAssignmentService)
    getAllAssignmentsController := controllers.NewGetAllAssignmentsController(getAllAssignmentsService)

    NfcCardAssignmentsRoutes(router, NfcCardAssignmentsHandlers{
        create:        createAssignmentController,
        getByID:      getAssignmentByIDController,
        getByUserID:  getAssignmentsByUserIDController,
        getByCardUID: getAssignmentByCardUIDController,
        update:       updateAssignmentController,
        deactivate:   deactivateAssignmentController,
        getAll:       getAllAssignmentsController,
    })
}