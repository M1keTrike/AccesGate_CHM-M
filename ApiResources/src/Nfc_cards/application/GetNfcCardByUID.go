package application

import (
	"api_resources/src/Nfc_cards/domain"
	"api_resources/src/Nfc_cards/domain/entities"
)

type GetNfcCardByUID struct {
	repo domain.NfcCardRepository
}

func NewGetNfcCardByUID(repo domain.NfcCardRepository) *GetNfcCardByUID {
	return &GetNfcCardByUID{repo: repo}
}

func (uc *GetNfcCardByUID) Execute(uid string) (entities.NfcCard, error) {
	return uc.repo.GetNfcCardByUID(uid)
}
