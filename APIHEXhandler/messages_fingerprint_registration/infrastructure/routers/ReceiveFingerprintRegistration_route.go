package routers

import (
    "github.com/M1keTrike/EventDriven/messages_fingerprint_registration/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

func AttachReceiveFingerprintRegistrationRoutes(r *gin.Engine, receiveFingerprintRegistrationController *controllers.ReceiveFingerprintRegistrationController) {
    fingerprint := r.Group("/fingerprint")
    {
        fingerprint.POST("/receive_registration", receiveFingerprintRegistrationController.Handle)
    }
}