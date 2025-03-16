package application

import "github.com/M1keTrike/EventDriven/services/websocket/domain/repositories"

type SendVAccessService struct {
	messageBus repositories.IMessageBusVAccess
}

func NewSendVAccessService(messageBus repositories.IMessageBusVAccess) SendVAccessService {
	return SendVAccessService{messageBus: messageBus}
}

func (s *SendVAccessService) Execute(msg []byte) error {

	return s.messageBus.Send(msg)
}
