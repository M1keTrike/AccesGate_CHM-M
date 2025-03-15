package application

import (
	"api_resources/src/Nfc_cards/domain"
	"api_resources/src/Nfc_cards/domain/entities"
)

type GetAllNfcCards struct {
	repo domain.NfcCardRepository
}

func NewGetAllNfcCards(repo domain.NfcCardRepository) *GetAllNfcCards {
	return &GetAllNfcCards{repo: repo}
}

func (uc *GetAllNfcCards) Execute() ([]entities.NfcCard, error) {
	cards, err := uc.repo.GetAllNfcCards()
	if err != nil {
		return nil, err
	}
	return cards, nil
}
