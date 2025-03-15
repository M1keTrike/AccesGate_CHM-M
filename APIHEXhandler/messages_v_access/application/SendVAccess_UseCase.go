package application

import "github.com/M1keTrike/EventDriven/services/websocket/application"

type SendVAccessUseCase struct {
	sa_s application.SendVAccessService
}

func NewSendVAccessUseCase(sa_s application.SendVAccessService) *SendVAccessUseCase {
	return &SendVAccessUseCase{sa_s: sa_s}
}

func (sa_uc SendVAccessUseCase) Execute(msg []byte) error {
	sa_uc.sa_s.Execute(msg)
	return nil
}
