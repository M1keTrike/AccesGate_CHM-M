package application

import "github.com/M1keTrike/EventDriven/services/websocket/domain/repositories"

type SendAlarmUseCase struct {
    messageBus repositories.IMessageBusAlarm
}

func NewSendAlarmUseCase(messageBus repositories.IMessageBusAlarm) SendAlarmUseCase {
    return SendAlarmUseCase{messageBus: messageBus}
}

func (s *SendAlarmUseCase) Execute(msg []byte) error {
    return s.messageBus.Send(msg)
}