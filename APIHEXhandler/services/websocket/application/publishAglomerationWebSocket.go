package application

import "github.com/M1keTrike/EventDriven/services/websocket/domain/repositories"

type SendAglomerationService struct {
	messageBus repositories.IMessageBusAglomeration
}

func NewSendAglomerationService(messageBus repositories.IMessageBusAglomeration) SendAglomerationService {
	return SendAglomerationService{messageBus: messageBus}
}

func (s *SendAglomerationService) Execute(msg []byte) error {

	return s.messageBus.Send(msg)
}