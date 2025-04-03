package application

import (
	"fmt"

	"github.com/M1keTrike/EventDriven/messages_nfc/domain"
)

type RecieveNFCMessageUseCase struct {
	sm_uc *SendNFCMessageUseCase
}

func NewRecieveNFCMessageUseCase(sm_uc *SendNFCMessageUseCase) RecieveNFCMessageUseCase {
	return RecieveNFCMessageUseCase{sm_uc: sm_uc}
}

func (r *RecieveNFCMessageUseCase) Execute(msg string) error {

	message := domain.NewMessage(msg)

	fmt.Println(message)

	ofJSON, err := domain.ToJSON(message)

	if err != nil {
		return err
	}
	r.sm_uc.Execute(ofJSON)

	return nil
}
