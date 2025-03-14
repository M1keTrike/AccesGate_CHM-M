package application

import "github.com/M1keTrike/EventDriven/messages_aglomeration/domain"

type ReceiveAglomerationUseCase struct {
	sa_uc *SendAglomerationUseCase
}

func NewRecieveAglomerationUseCase(sa_uc *SendAglomerationUseCase) ReceiveAglomerationUseCase {
	return ReceiveAglomerationUseCase{sa_uc: sa_uc}
}

func (r *ReceiveAglomerationUseCase) Execute(msg string) error {

	message := domain.NewMessage(msg)
	msJSON, err := domain.ToJSON(message)
	if err != nil {
		return err
	}

	r.sa_uc.Execute(msJSON)
	return nil
}
