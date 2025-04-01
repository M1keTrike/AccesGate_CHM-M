package domain

import "api_resources/src/Nfc_cards/domain/entities"

type NfcCardRepository interface {
	CreateNfcCard(card *entities.NfcCard) error
	GetNfcCardByUID(uid string) (entities.NfcCard, error)
	DeleteNfcCard(uid string) error
	GetAllNfcCards() ([]entities.NfcCard, error)
}
