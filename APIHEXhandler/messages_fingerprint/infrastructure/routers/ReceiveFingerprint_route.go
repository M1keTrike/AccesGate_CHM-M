package routers

import (
	"github.com/M1keTrike/EventDriven/messages_fingerprint/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func AttachReceiveFingerprintRoutes(r *gin.Engine, receiveFingerprintController *controllers.ReceiveFingerprintController) {
	fingerprint := r.Group("/fingerprint")
	{
		fingerprint.POST("/receive_fingerprint", receiveFingerprintController.Handle)
	}
}