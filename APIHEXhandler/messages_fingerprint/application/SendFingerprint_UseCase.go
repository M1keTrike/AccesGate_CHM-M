package application

import "github.com/M1keTrike/EventDriven/services/websocket/domain/repositories"

type SendFingerprintUseCase struct {
	messageBus repositories.IMessageBusFingerprint
}

func NewSendFingerprintUseCase(messageBus repositories.IMessageBusFingerprint) SendFingerprintUseCase {
	return SendFingerprintUseCase{messageBus: messageBus}
}

func (s *SendFingerprintUseCase) Execute(msg []byte) error {
	return s.messageBus.Send(msg)
}