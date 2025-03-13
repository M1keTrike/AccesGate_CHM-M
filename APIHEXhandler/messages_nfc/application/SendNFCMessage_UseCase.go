package application

import "github.com/M1keTrike/EventDriven/services/websocket/application"

type SendNFCMessageUseCase struct {
	sm_s *application.SendWSNFCService
}

func NewSendMessageUseCase(sm_s *application.SendWSNFCService) *SendNFCMessageUseCase {
	return &SendNFCMessageUseCase{sm_s: sm_s}
}

func (s *SendNFCMessageUseCase) Execute(msg []byte) error {

	s.sm_s.Execute(msg)
	return nil
}
