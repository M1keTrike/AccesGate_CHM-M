package application

import "github.com/M1keTrike/EventDriven/messages_aglomeration/domain"

type ReceiveVAccessUseCase struct {
	sa_uc *SendVAccessUseCase
}

func NewRecieveVAccessUseCase(sa_uc *SendVAccessUseCase) ReceiveVAccessUseCase {
	return ReceiveVAccessUseCase{sa_uc: sa_uc}
}

func (r *ReceiveVAccessUseCase) Execute(msg string) error {

	message := domain.NewMessage(msg)
	msJSON, err := domain.ToJSON(message)
	if err != nil {
		return err
	}

	r.sa_uc.Execute(msJSON)
	return nil
}
