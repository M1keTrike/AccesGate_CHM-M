package application

import (
	"api_resources/src/Nfc_cards/domain"
	"api_resources/src/Nfc_cards/domain/entities"
)

type CreateNfcCard struct {
	repo domain.NfcCardRepository
}

func NewCreateNfcCard(repo domain.NfcCardRepository) *CreateNfcCard {
	return &CreateNfcCard{repo: repo}
}

func (uc *CreateNfcCard) Execute(card *entities.NfcCard) error {
	err := uc.repo.CreateNfcCard(card)
	if err != nil {
		return err
	}
	return nil
}
