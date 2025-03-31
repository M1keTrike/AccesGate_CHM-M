package infrastructure

import (
    "database/sql"
    "fmt"
    "time"
    "api_resources/src/core"
    "api_resources/src/Events/domain/entities"
)

type PostgreSQL struct {
    conn *core.Conn_PostgreSQL
}

func NewPostgreSQL() *PostgreSQL {
    conn := core.GetDBPool()
    return &PostgreSQL{conn: conn}
}

func (pg *PostgreSQL) CreateEvent(event entities.Event) error {
    query := `
        INSERT INTO events (name, description, start_time, end_time, created_by, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

    event.CreatedAt = time.Now()
    err := pg.conn.DB.QueryRow(
        query,
        event.Name,
        event.Description,
        event.StartTime.Format("2006-01-02 15:04:05"),
        event.EndTime.Format("2006-01-02 15:04:05"),
        event.CreatedBy,
        event.CreatedAt.Format("2006-01-02 15:04:05"),
    ).Scan(&event.ID)

    if err != nil {
        return fmt.Errorf("error creating event: %v", err)
    }
    return nil
}

func (pg *PostgreSQL) GetEventByID(id int) (entities.Event, error) {
    var event entities.Event
    var startTimeStr, endTimeStr, createdAtStr string

    query := `SELECT id, name, description, start_time, end_time, created_by, created_at 
              FROM events WHERE id = $1`

    err := pg.conn.DB.QueryRow(query, id).Scan(
        &event.ID,
        &event.Name,
        &event.Description,
        &startTimeStr,
        &endTimeStr,
        &event.CreatedBy,
        &createdAtStr,
    )

    if err == sql.ErrNoRows {
        return event, fmt.Errorf("event not found")
    }
    if err != nil {
        return event, fmt.Errorf("error getting event: %v", err)
    }

    event.StartTime, _ = time.Parse("2006-01-02 15:04:05", startTimeStr)
    event.EndTime, _ = time.Parse("2006-01-02 15:04:05", endTimeStr)
    event.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtStr)

    return event, nil
}

func (pg *PostgreSQL) GetAllEvents() ([]entities.Event, error) {
    query := `SELECT id, name, description, start_time, end_time, created_by, created_at 
              FROM events`

    rows, err := pg.conn.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error getting events: %v", err)
    }
    defer rows.Close()

    var events []entities.Event
    for rows.Next() {
        var event entities.Event
        var startTimeStr, endTimeStr, createdAtStr string

        err := rows.Scan(
            &event.ID,
            &event.Name,
            &event.Description,
            &startTimeStr,
            &endTimeStr,
            &event.CreatedBy,
            &createdAtStr,
        )
        if err != nil {
            return nil, fmt.Errorf("error scanning event: %v", err)
        }

        event.StartTime, _ = time.Parse("2006-01-02 15:04:05", startTimeStr)
        event.EndTime, _ = time.Parse("2006-01-02 15:04:05", endTimeStr)
        event.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtStr)

        events = append(events, event)
    }

    return events, nil
}

func (pg *PostgreSQL) UpdateEvent(event entities.Event) error {
    query := `
        UPDATE events 
        SET name = $1, description = $2, start_time = $3, end_time = $4
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
        return fmt.Errorf("error updating event: %v", err)
    }
    if result == nil {
        return fmt.Errorf("event not found")
    }
    return nil
}

func (pg *PostgreSQL) DeleteEvent(id int) error {
    query := "DELETE FROM events WHERE id = $1"
    result, err := pg.conn.ExecutePreparedQuery(query, id)
    if err != nil {
        return fmt.Errorf("error deleting event: %v", err)
    }
    if result == nil {
        return fmt.Errorf("event not found")
    }
    return nil
}