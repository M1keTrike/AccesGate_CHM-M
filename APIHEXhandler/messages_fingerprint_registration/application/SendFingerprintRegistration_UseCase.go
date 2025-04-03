package application

import "github.com/M1keTrike/EventDriven/services/websocket/domain/repositories"

type SendFingerprintRegistrationUseCase struct {
    messageBus repositories.IMessageBusFingerprintRegistration
}

func NewSendFingerprintRegistrationUseCase(messageBus repositories.IMessageBusFingerprintRegistration) SendFingerprintRegistrationUseCase {
    return SendFingerprintRegistrationUseCase{messageBus: messageBus}
}

func (s *SendFingerprintRegistrationUseCase) Execute(msg []byte) error {
    return s.messageBus.Send(msg)
}