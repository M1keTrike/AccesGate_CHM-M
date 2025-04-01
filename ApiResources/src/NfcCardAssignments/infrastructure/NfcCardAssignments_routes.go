package infrastructure

import (
    "api_resources/src/NfcCardAssignments/infrastructure/controllers"
    "api_resources/src/core"
    "github.com/gin-gonic/gin"
)

type NfcCardAssignmentsHandlers struct {
    create            *controllers.CreateAssignmentController
    getByID          *controllers.GetAssignmentByIDController
    getByUserID      *controllers.GetAssignmentsByUserIDController
    getByCardUID     *controllers.GetAssignmentByCardUIDController
    update           *controllers.UpdateAssignmentController
    deactivate       *controllers.DeactivateAssignmentController
    getAll           *controllers.GetAllAssignmentsController
}

func NfcCardAssignmentsRoutes(router *gin.Engine, handlers NfcCardAssignmentsHandlers) {
    protected := router.Group("/api/nfc-assignments")
    protected.Use(core.AuthMiddleware())
    {
        protected.POST("", handlers.create.Handle)
        protected.GET("/:id", handlers.getByID.Handle)
        protected.GET("/user/:userId", handlers.getByUserID.Handle)
        protected.GET("/card/:cardUid", handlers.getByCardUID.Handle)
        protected.PUT("/:id", handlers.update.Handle)
        protected.PUT("/:id/deactivate", handlers.deactivate.Handle)
        protected.GET("", handlers.getAll.Handle)
    }
}