package main

import (
	"log"
	"os"

	aglomerationDeps "github.com/M1keTrike/EventDriven/messages_aglomeration/dependencies"
	alarmDeps "github.com/M1keTrike/EventDriven/messages_alarm/dependencies"
	fingerprintDeps "github.com/M1keTrike/EventDriven/messages_fingerprint/dependencies"
	fingerprintRegistrationDeps "github.com/M1keTrike/EventDriven/messages_fingerprint_registration/dependencies"
	vaccessDeps "github.com/M1keTrike/EventDriven/messages_v_access/dependencies"

	"github.com/M1keTrike/EventDriven/messages_nfc/dependencies"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	r := gin.Default()

	messageNFCDependencies := dependencies.NewMessageDependencies()
	messageNFCDependencies.Execute(r)

	messageAglomerationDependencies := aglomerationDeps.NewMessageAgloemrationDependencies()
	messageAglomerationDependencies.Execute(r)

	messageVAccessDependencies := vaccessDeps.NewMessageAgloemrationDependencies()
	messageVAccessDependencies.Execute(r)

	messageFingerprintDependencies := fingerprintDeps.NewMessageFingerprintDependencies()
	messageFingerprintDependencies.Execute(r)

	messageFingerprintRegistrationDeps := fingerprintRegistrationDeps.NewMessageFingerprintRegistrationDependencies()
	messageFingerprintRegistrationDeps.Execute(r)

	messageAlarmDeps := alarmDeps.NewMessageAlarmDependencies()
	messageAlarmDeps.Execute(r)

	r.Run(":" + PORT)
}
