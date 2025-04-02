package infrastructure

import (
	"api_resources/src/Events/domain"
	"api_resources/src/Events/domain/entities"
	"api_resources/src/core"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type PostgreSQL struct {
	conn *core.Conn_PostgreSQL
}

func NewPostgreSQLEvents() domain.EventRepository {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("[DB INIT - EVENTS] Error en pool de conexiones: %v", conn.Err)
	}
	return &PostgreSQL{conn: conn}
}

// CreateEvent crea un nuevo evento
func (pg *PostgreSQL) CreateEvent(event *entities.Event) error {
	query := `
		INSERT INTO events (name, description, start_time, end_time, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	event.CreatedAt = time.Now().UTC() // Use UTC time
	err := pg.conn.DB.QueryRow(
		query,
		event.Name,
		event.Description,
		event.StartTime.UTC().Format("2006-01-02 15:04:05"),
		event.EndTime.UTC().Format("2006-01-02 15:04:05"),
		event.CreatedBy,
		event.CreatedAt.Format("2006-01-02 15:04:05"),
	).Scan(&event.ID)

	if err != nil {
		log.Printf("[CreateEvent] Error al crear evento: %v", err)
		return fmt.Errorf("no se pudo crear el evento")
	}
	log.Printf("[CreateEvent] Evento '%s' creado con ID %d", event.Name, event.ID)
	return nil
}

// Modify the GetAllEvents function
func (pg *PostgreSQL) GetAllEvents() ([]entities.Event, error) {
	query := `SELECT id, name, description, start_time, end_time, created_by, created_at FROM events`
	rows, err := pg.conn.DB.Query(query)
	if err != nil {
		log.Printf("[GetAllEvents] Error: %v", err)
		return nil, fmt.Errorf("no se pudieron obtener los eventos")
	}
	defer rows.Close()

	var events []entities.Event
	for rows.Next() {
		var event entities.Event
		var startTime, endTime, createdAt time.Time

		if err := rows.Scan(
			&event.ID, &event.Name, &event.Description,
			&startTime, &endTime, &event.CreatedBy, &createdAt,
		); err != nil {
			log.Printf("[GetAllEvents] Error al escanear fila: %v", err)
			continue
		}

		event.StartTime = startTime
		event.EndTime = endTime
		event.CreatedAt = createdAt

		events = append(events, event)
	}

	return events, nil
}

// Modify GetEventByID similarly
func (pg *PostgreSQL) GetEventByID(id int) (*entities.Event, error) {
	query := `SELECT id, name, description, start_time, end_time, created_by, created_at FROM events WHERE id = $1`
	var event entities.Event
	var startTime, endTime, createdAt time.Time

	err := pg.conn.DB.QueryRow(query, id).Scan(
		&event.ID, &event.Name, &event.Description,
		&startTime, &endTime, &event.CreatedBy, &createdAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("evento no encontrado")
	}
	if err != nil {
		log.Printf("[GetEventByID] Error: %v", err)
		return nil, fmt.Errorf("error al buscar el evento")
	}

	event.StartTime = startTime
	event.EndTime = endTime
	event.CreatedAt = createdAt

	return &event, nil
}

// Update GetEventsByCreator similarly
func (pg *PostgreSQL) GetEventsByCreator(userID int) ([]entities.Event, error) {
	query := `SELECT id, name, description, start_time, end_time, created_by, created_at FROM events WHERE created_by = $1`
	rows, err := pg.conn.DB.Query(query, userID)
	if err != nil {
		log.Printf("[GetEventsByCreator] Error: %v", err)
		return nil, fmt.Errorf("no se pudieron obtener los eventos")
	}
	defer rows.Close()

	var events []entities.Event
	for rows.Next() {
		var event entities.Event
		var startTime, endTime, createdAt time.Time

		if err := rows.Scan(
			&event.ID, &event.Name, &event.Description,
			&startTime, &endTime, &event.CreatedBy, &createdAt,
		); err != nil {
			log.Printf("[GetEventsByCreator] Error al escanear fila: %v", err)
			continue
		}

		event.StartTime = startTime
		event.EndTime = endTime
		event.CreatedAt = createdAt

		events = append(events, event)
	}

	return events, nil
}

// UpdateEvent actualiza un evento existente
func (pg *PostgreSQL) UpdateEvent(event *entities.Event) error {
	query := `
		UPDATE events SET name = $1, description = $2, start_time = $3, end_time = $4
		WHERE id = $5`
	result, err := pg.conn.ExecutePreparedQuery(
		query,
		event.Name,
		event.Description,
		event.StartTime.Format("2006-01-02 15:04:05"),
		event.EndTime.Format("2006-01-02 15:04:05"),
		event.ID,
	)
	if err != nil {
		log.Printf("[UpdateEvent] Error al actualizar evento ID %d: %v", event.ID, err)
		return fmt.Errorf("no se pudo actualizar el evento")
	}
	if result == nil {
		return fmt.Errorf("evento no encontrado")
	}
	log.Printf("[UpdateEvent] Evento ID %d actualizado", event.ID)
	return nil
}

// DeleteEvent elimina un evento por ID
func (pg *PostgreSQL) DeleteEvent(id int) error {
	query := `DELETE FROM events WHERE id = $1`
	result, err := pg.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("[DeleteEvent] Error al eliminar evento ID %d: %v", id, err)
		return fmt.Errorf("no se pudo eliminar el evento")
	}
	if result == nil {
		return fmt.Errorf("evento no encontrado")
	}
	log.Printf("[DeleteEvent] Evento ID %d eliminado", id)
	return nil
}
