package application

import "github.com/M1keTrike/EventDriven/services/websocket/application"

type SendAglomerationUseCase struct {
	sa_s application.SendAglomerationService
}

func NewSendAglomerationUseCase(sa_s application.SendAglomerationService) *SendAglomerationUseCase {
	return &SendAglomerationUseCase{sa_s: sa_s}
}

func (sa_uc SendAglomerationUseCase) Execute(msg []byte) error {
	sa_uc.sa_s.Execute(msg)
	return nil
}
