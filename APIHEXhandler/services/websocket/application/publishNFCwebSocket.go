package application

import "github.com/M1keTrike/EventDriven/services/websocket/domain/repositories"

type SendWSNFCService struct {
	messageBus repositories.IMessageBusNFC
}


func NewSendWSNFCService(messageBus repositories.IMessageBusNFC) *SendWSNFCService {
	return &SendWSNFCService{messageBus: messageBus}
}

func (s *SendWSNFCService) Execute(msg []byte) error {

	return s.messageBus.Send(msg)
}
