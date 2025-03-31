package infraestructure

import (
    "api_resources/src/core"
    "api_resources/src/EventAttendees/domain/entities"
    _"database/sql"
	"log"
    "fmt"
    "time"
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

func (pg *PostgreSQL) RegisterAttendee(attendee *entities.EventAttendee) error {
    query := `
        INSERT INTO event_attendees (user_id, event_id, registered_at) 
        VALUES ($1, $2, $3) 
        RETURNING id`
    
    attendee.RegisteredAt = time.Now()
    err := pg.conn.DB.QueryRow(query, attendee.UserID, attendee.EventID, attendee.RegisteredAt).Scan(&attendee.ID)
    if err != nil {
        return fmt.Errorf("error registering attendee: %v", err)
    }
    return nil
}

func (pg *PostgreSQL) RemoveAttendee(eventID, userID int) error {
    query := "DELETE FROM event_attendees WHERE event_id = $1 AND user_id = $2"
    _, err := pg.conn.ExecutePreparedQuery(query, eventID, userID)
    if err != nil {
        return fmt.Errorf("error removing attendee: %v", err)
    }
    return nil
}

func (pg *PostgreSQL) GetEventAttendees(eventID int) ([]entities.EventAttendee, error) {
    attendees := []entities.EventAttendee{}
    query := "SELECT id, user_id, event_id, registered_at FROM event_attendees WHERE event_id = $1"
    
    rows, err := pg.conn.DB.Query(query, eventID)
    if err != nil {
        return nil, fmt.Errorf("error getting event attendees: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var attendee entities.EventAttendee
        if err := rows.Scan(&attendee.ID, &attendee.UserID, &attendee.EventID, &attendee.RegisteredAt); err != nil {
            return nil, fmt.Errorf("error scanning attendee: %v", err)
        }
        attendees = append(attendees, attendee)
    }

    return attendees, nil
}

func (pg *PostgreSQL) GetUserEvents(userID int) ([]entities.EventAttendee, error) {
    attendees := []entities.EventAttendee{}
    query := "SELECT id, user_id, event_id, registered_at FROM event_attendees WHERE user_id = $1"
    
    rows, err := pg.conn.DB.Query(query, userID)
    if err != nil {
        return nil, fmt.Errorf("error getting user events: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var attendee entities.EventAttendee
        if err := rows.Scan(&attendee.ID, &attendee.UserID, &attendee.EventID, &attendee.RegisteredAt); err != nil {
            return nil, fmt.Errorf("error scanning attendee: %v", err)
        }
        attendees = append(attendees, attendee)
    }

    return attendees, nil
}

func (pg *PostgreSQL) IsUserRegistered(eventID, userID int) (bool, error) {
    var exists bool
    query := "SELECT EXISTS(SELECT 1 FROM event_attendees WHERE event_id = $1 AND user_id = $2)"
    
    err := pg.conn.DB.QueryRow(query, eventID, userID).Scan(&exists)
    if err != nil {
        return false, fmt.Errorf("error checking registration: %v", err)
    }
    
    return exists, nil
}