package application

import (
	"api_resources/src/Nfc_cards/domain"
)

type DeleteNfcCard struct {
	repo domain.NfcCardRepository
}

func NewDeleteNfcCard(repo domain.NfcCardRepository) *DeleteNfcCard {
	return &DeleteNfcCard{repo: repo}
}

func (uc *DeleteNfcCard) Execute(uid string) error {
	return uc.repo.DeleteNfcCard(uid)
}
