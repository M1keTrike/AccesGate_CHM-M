package infraestructure

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    "api_resources/src/core"
    "api_resources/src/EventAttendees/domain/entities"
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
        INSERT INTO event_attendees (user_id, event_id, registered_at, attended) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id`
    
    attendee.RegisteredAt = time.Now()
    attendee.Attended = false // Default value for new registrations
    formattedTime := attendee.RegisteredAt.Format("2006-01-02 15:04:05")
    err := pg.conn.DB.QueryRow(query, attendee.UserID, attendee.EventID, formattedTime, attendee.Attended).Scan(&attendee.ID)
    if err != nil {
        if err == sql.ErrNoRows {
            return fmt.Errorf("error registering attendee: no rows affected")
        }
        return fmt.Errorf("error registering attendee: %v", err)
    }
    return nil
}

func (pg *PostgreSQL) RemoveAttendee(eventID, userID int) error {
    query := "DELETE FROM event_attendees WHERE event_id = $1 AND user_id = $2"
    result, err := pg.conn.ExecutePreparedQuery(query, eventID, userID)
    if err != nil {
        return fmt.Errorf("error removing attendee: %v", err)
    }
    if result == nil {
        return fmt.Errorf("attendee not found for event %d and user %d", eventID, userID)
    }
    return nil
}

func parseDateTime(dateStr string) (time.Time, error) {
    layouts := []string{
        "2006-01-02 15:04:05",
        time.RFC3339,
        "2006-01-02T15:04:05Z",
    }

    for _, layout := range layouts {
        if t, err := time.Parse(layout, dateStr); err == nil {
            return t, nil
        }
    }
    return time.Time{}, fmt.Errorf("could not parse date: %s", dateStr)
}

func (pg *PostgreSQL) GetEventAttendees(eventID int) ([]entities.EventAttendee, error) {
    attendees := []entities.EventAttendee{}
    // Use COALESCE para manejar valores NULL
    query := "SELECT id, user_id, event_id, registered_at, COALESCE(attended, false) as attended FROM event_attendees WHERE event_id = $1"
    
    rows, err := pg.conn.DB.Query(query, eventID)
    if err != nil {
        return nil, fmt.Errorf("error getting event attendees: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var attendee entities.EventAttendee
        var registeredAtStr string
        if err := rows.Scan(&attendee.ID, &attendee.UserID, &attendee.EventID, &registeredAtStr, &attendee.Attended); err != nil {
            return nil, fmt.Errorf("error scanning attendee: %v", err)
        }
        parsedTime, err := parseDateTime(registeredAtStr)
        if err != nil {
            return nil, fmt.Errorf("error parsing registration date: %v", err)
        }
        attendee.RegisteredAt = parsedTime
        attendees = append(attendees, attendee)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating event attendees: %v", err)
    }

    return attendees, nil
}

func (pg *PostgreSQL) GetUserEvents(userID int) ([]entities.EventAttendee, error) {
    attendees := []entities.EventAttendee{}
    query := "SELECT id, user_id, event_id, registered_at, COALESCE(attended, false) as attended FROM event_attendees WHERE user_id = $1"
    
    rows, err := pg.conn.DB.Query(query, userID)
    if err != nil {
        return nil, fmt.Errorf("error getting user events: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var attendee entities.EventAttendee
        var registeredAtStr string
        if err := rows.Scan(&attendee.ID, &attendee.UserID, &attendee.EventID, &registeredAtStr, &attendee.Attended); err != nil {
            return nil, fmt.Errorf("error scanning attendee: %v", err)
        }
        parsedTime, err := parseDateTime(registeredAtStr)
        if err != nil {
            return nil, fmt.Errorf("error parsing registration date: %v", err)
        }
        attendee.RegisteredAt = parsedTime
        attendees = append(attendees, attendee)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating user events: %v", err)
    }

    return attendees, nil
}

func (pg *PostgreSQL) IsUserRegistered(eventID, userID int) (bool, error) {
    var exists bool
    query := "SELECT EXISTS(SELECT 1 FROM event_attendees WHERE event_id = $1 AND user_id = $2)"
    
    err := pg.conn.DB.QueryRow(query, eventID, userID).Scan(&exists)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil
        }
        return false, fmt.Errorf("error checking registration: %v", err)
    }
    
    return exists, nil
}

func (pg *PostgreSQL) UpdateAttendanceStatus(eventID, userID int, attended bool) error {
    query := `
        UPDATE event_attendees 
        SET attended = $1 
        WHERE event_id = $2 AND user_id = $3
    `
    
    result, err := pg.conn.DB.Exec(query, attended, eventID, userID)
    if err != nil {
        return fmt.Errorf("error updating attendance status: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %v", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no attendee found with eventID %d and userID %d", eventID, userID)
    }

    return nil
}