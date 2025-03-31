package infraestructure

import (
	"api_resources/src/Events/domain/entities"
	"api_resources/src/core"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type EventPostgreSQL struct {
	conn *core.Conn_PostgreSQL
}

func NewEventPostgreSQL() *EventPostgreSQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &EventPostgreSQL{conn: conn}
}

func (pg *EventPostgreSQL) CreateEvent(event entities.Event) error {
	query := `INSERT INTO events (name, description, start_time, end_time, created_by, created_at)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return pg.conn.DB.QueryRow(query, event.Name, event.Description, event.StartTime, event.EndTime, event.CreatedBy, event.CreatedAt).Scan(&event.ID)
}

func (pg *EventPostgreSQL) GetEventByID(id int) (entities.Event, error) {
	event := entities.Event{}
	var createdAt, startTime, endTime string

	query := `SELECT id, name, description, start_time, end_time, created_by, created_at FROM events WHERE id = $1`
	row := pg.conn.DB.QueryRow(query, id)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &startTime, &endTime, &event.CreatedBy, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return event, fmt.Errorf("evento con ID %d no encontrado", id)
		}
		return event, fmt.Errorf("error al obtener evento: %v", err)
	}

	event.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	event.StartTime, _ = time.Parse(time.RFC3339, startTime)
	event.EndTime, _ = time.Parse(time.RFC3339, endTime)

	return event, nil
}

func (pg *EventPostgreSQL) GetAllEvents() ([]entities.Event, error) {
	query := `SELECT id, name, description, start_time, end_time, created_by, created_at FROM events`
	rows, err := pg.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener eventos: %v", err)
	}
	defer rows.Close()

	var events []entities.Event

	for rows.Next() {
		var event entities.Event
		var createdAt, startTime, endTime string

		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &startTime, &endTime, &event.CreatedBy, &createdAt); err != nil {
			return nil, fmt.Errorf("error al escanear eventos: %v", err)
		}

		event.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		event.StartTime, _ = time.Parse(time.RFC3339, startTime)
		event.EndTime, _ = time.Parse(time.RFC3339, endTime)

		events = append(events, event)
	}

	return events, nil
}

func (pg *EventPostgreSQL) UpdateEvent(event entities.Event) error {
	query := `UPDATE events SET name=$1, description=$2, start_time=$3, end_time=$4 WHERE id=$5`
	_, err := pg.conn.ExecutePreparedQuery(query, event.Name, event.Description, event.StartTime, event.EndTime, event.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar evento: %v", err)
	}
	return nil
}

func (pg *EventPostgreSQL) DeleteEvent(id int) error {
	query := `DELETE FROM events WHERE id=$1`
	_, err := pg.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar evento: %v", err)
	}
	return nil
}
