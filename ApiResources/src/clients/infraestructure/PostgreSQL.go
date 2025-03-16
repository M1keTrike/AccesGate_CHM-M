package infraestructure

import (
	"api_resources/src/clients/domain/entities"
	"database/sql"
	"fmt"
	"log"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewPostgreSQL() *PostgreSQL {
	connStr := "user=username dbname=yourdb sslmode=disable password=yourpassword"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	return &PostgreSQL{db: db}
}

func (pg *PostgreSQL) CreateClient(client *entities.Client) error {
	query := "INSERT INTO clients (worker_id, front_id, hardware_id) VALUES ($1, $2, $3) RETURNING client_id"
	err := pg.db.QueryRow(query, client.WorkerID, client.FrontID, client.HardwareID).Scan(&client.ClientID)
	if err != nil {
		return fmt.Errorf("error creating client: %v", err)
	}
	return nil
}

func (pg *PostgreSQL) GetAllClients() ([]entities.Client, error) {
	query := "SELECT client_id, worker_id, front_id, hardware_id FROM clients"
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching clients: %v", err)
	}
	defer rows.Close()

	var clients []entities.Client
	for rows.Next() {
		var client entities.Client
		if err := rows.Scan(&client.ClientID, &client.WorkerID, &client.FrontID, &client.HardwareID); err != nil {
			return nil, fmt.Errorf("error scanning client: %v", err)
		}
		clients = append(clients, client)
	}
	return clients, nil
}
