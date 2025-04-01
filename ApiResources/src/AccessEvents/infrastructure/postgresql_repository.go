package infrastructure

import (
	"api_resources/src/AccessEvents/domain"
	"api_resources/src/AccessEvents/domain/entities"
	"api_resources/src/core"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type PostgreSQL struct {
	conn *core.Conn_PostgreSQL
}

func NewPostgreSQLAccessEvents() domain.AccessEventRepository {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("[DB INIT - ACCESS_EVENTS] Error en pool de conexiones: %v", conn.Err)
	}
	return &PostgreSQL{conn: conn}
}

func (pg *PostgreSQL) Create(event *entities.AccessEvent) error {
	query := `
		INSERT INTO access_events (user_id, front_id, device_id, status, timestamp)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	event.Timestamp = time.Now()
	err := pg.conn.DB.QueryRow(
		query,
		event.UserID,
		event.FrontID,
		event.DeviceID,
		event.Status,
		event.Timestamp.Format("2006-01-02 15:04:05"),
	).Scan(&event.ID)

	if err != nil {
		log.Printf("[CreateAccessEvent] Error al crear evento: %v", err)
		return fmt.Errorf("no se pudo crear el evento")
	}
	return nil
}

func (pg *PostgreSQL) GetAll() ([]entities.AccessEvent, error) {
	query := `SELECT id, user_id, front_id, device_id, status, timestamp FROM access_events`
	rows, err := pg.conn.DB.Query(query)
	if err != nil {
		log.Printf("[GetAllAccessEvents] Error: %v", err)
		return nil, fmt.Errorf("no se pudieron obtener los eventos de acceso")
	}
	defer rows.Close()

	var events []entities.AccessEvent
	for rows.Next() {
		var event entities.AccessEvent
		var timestampStr string

		if err := rows.Scan(&event.ID, &event.UserID, &event.FrontID, &event.DeviceID, &event.Status, &timestampStr); err != nil {
			log.Printf("[GetAllAccessEvents] Error al escanear fila: %v", err)
			continue
		}
		event.Timestamp, _ = time.Parse("2006-01-02 15:04:05", timestampStr)
		events = append(events, event)
	}

	return events, nil
}

func (pg *PostgreSQL) GetByID(id int) (*entities.AccessEvent, error) {
	query := `SELECT id, user_id, front_id, device_id, status, timestamp FROM access_events WHERE id = $1`
	var event entities.AccessEvent
	var timestampStr string

	err := pg.conn.DB.QueryRow(query, id).Scan(
		&event.ID, &event.UserID, &event.FrontID, &event.DeviceID, &event.Status, &timestampStr,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("evento no encontrado")
	}
	if err != nil {
		log.Printf("[GetAccessEventByID] Error: %v", err)
		return nil, fmt.Errorf("error al obtener el evento")
	}

	event.Timestamp, _ = time.Parse("2006-01-02 15:04:05", timestampStr)
	return &event, nil
}

func (pg *PostgreSQL) GetByUser(userID int) ([]entities.AccessEvent, error) {
	query := `SELECT id, user_id, front_id, device_id, status, timestamp FROM access_events WHERE user_id = $1`
	return pg.getByField(query, userID)
}

func (pg *PostgreSQL) GetByDevice(deviceID int) ([]entities.AccessEvent, error) {
	query := `SELECT id, user_id, front_id, device_id, status, timestamp FROM access_events WHERE device_id = $1`
	return pg.getByField(query, deviceID)
}

func (pg *PostgreSQL) GetByFront(frontID int) ([]entities.AccessEvent, error) {
	query := `SELECT id, user_id, front_id, device_id, status, timestamp FROM access_events WHERE front_id = $1`
	return pg.getByField(query, frontID)
}

func (pg *PostgreSQL) getByField(query string, value interface{}) ([]entities.AccessEvent, error) {
	rows, err := pg.conn.DB.Query(query, value)
	if err != nil {
		log.Printf("[GetAccessEventsByField] Error: %v", err)
		return nil, fmt.Errorf("no se pudieron obtener los eventos de acceso")
	}
	defer rows.Close()

	var events []entities.AccessEvent
	for rows.Next() {
		var event entities.AccessEvent
		var timestampStr string

		if err := rows.Scan(&event.ID, &event.UserID, &event.FrontID, &event.DeviceID, &event.Status, &timestampStr); err != nil {
			log.Printf("[ScanAccessEvent] Error: %v", err)
			continue
		}
		event.Timestamp, _ = time.Parse("2006-01-02 15:04:05", timestampStr)
		events = append(events, event)
	}
	return events, nil
}

func (pg *PostgreSQL) Update(event *entities.AccessEvent) error {
	query := `
		UPDATE access_events
		SET user_id = $1, front_id = $2, device_id = $3, status = $4, timestamp = $5
		WHERE id = $6`
	result, err := pg.conn.ExecutePreparedQuery(
		query,
		event.UserID,
		event.FrontID,
		event.DeviceID,
		event.Status,
		event.Timestamp.Format("2006-01-02 15:04:05"),
		event.ID,
	)
	if err != nil {
		log.Printf("[UpdateAccessEvent] Error: %v", err)
		return fmt.Errorf("no se pudo actualizar el evento")
	}
	if result == nil {
		return fmt.Errorf("evento no encontrado")
	}
	return nil
}

func (pg *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM access_events WHERE id = $1`
	result, err := pg.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("[DeleteAccessEvent] Error: %v", err)
		return fmt.Errorf("no se pudo eliminar el evento")
	}
	if result == nil {
		return fmt.Errorf("evento no encontrado")
	}
	return nil
}
