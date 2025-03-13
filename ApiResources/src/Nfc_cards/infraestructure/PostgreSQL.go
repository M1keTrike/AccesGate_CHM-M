package infraestructure

import (
	"database/sql"
	"fmt"
	"log"

	"api_resources/src/Nfc_cards/domain/entities"
	"api_resources/src/core"
)

type PostgreSQL struct {
	conn *core.Conn_PostgreSQL
}

func NewPostgreSQL() *PostgreSQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &PostgreSQL{conn: conn}
}

func (pg *PostgreSQL) GetNfcCardByUID(cardUID string) (entities.NfcCard, error) {
	card := entities.NfcCard{}
	query := "SELECT card_uid FROM nfc_cards WHERE card_uid = $1"
	row := pg.conn.DB.QueryRow(query, cardUID)

	err := row.Scan(&card.CardUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return card, fmt.Errorf("tarjeta con UID %s no encontrada", cardUID)
		}
		return card, fmt.Errorf("error al obtener tarjeta: %v", err)
	}

	return card, nil
}

func (pg *PostgreSQL) CreateNfcCard(card *entities.NfcCard) error {
	query := "INSERT INTO nfc_cards (card_uid) VALUES ($1)"
	_, err := pg.conn.ExecutePreparedQuery(query, card.CardUID)
	if err != nil {
		return fmt.Errorf("error al crear tarjeta: %v", err)
	}
	return nil
}

func (pg *PostgreSQL) DeleteNfcCard(cardUID string) error {
	query := "DELETE FROM nfc_cards WHERE card_uid = $1"
	_, err := pg.conn.ExecutePreparedQuery(query, cardUID)
	if err != nil {
		return fmt.Errorf("error al eliminar tarjeta: %v", err)
	}
	return nil
}

func (pg *PostgreSQL) GetAllNfcCards() ([]entities.NfcCard, error) {
	cards := []entities.NfcCard{}
	query := "SELECT card_uid FROM nfc_cards"
	rows, err := pg.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener tarjetas: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		card := entities.NfcCard{}
		if err := rows.Scan(&card.CardUID); err != nil {
			return nil, fmt.Errorf("error al escanear tarjetas: %v", err)
		}
		cards = append(cards, card)
	}

	return cards, nil
}
