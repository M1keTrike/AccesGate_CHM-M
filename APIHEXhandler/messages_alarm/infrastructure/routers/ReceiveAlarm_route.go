package routers

import (
    "github.com/M1keTrike/EventDriven/messages_alarm/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

func AttachReceiveAlarmRoutes(r *gin.Engine, receiveAlarmController *controllers.ReceiveAlarmController) {
    alarm := r.Group("/alarm")
    {
        alarm.POST("/receive_alarm", receiveAlarmController.Handle)
    }
}