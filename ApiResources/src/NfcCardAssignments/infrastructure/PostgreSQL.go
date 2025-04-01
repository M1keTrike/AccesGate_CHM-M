package infrastructure

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    "api_resources/src/core"
    "api_resources/src/NfcCardAssignments/domain/entities"
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

func (pg *PostgreSQL) CreateAssignment(assignment *entities.NfcCardAssignment) error {
    query := `
        INSERT INTO nfc_card_assignments (user_id, card_uid, assigned_at, is_active)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

    assignment.AssignedAt = time.Now()
    assignment.IsActive = true

    err := pg.conn.DB.QueryRow(
        query,
        assignment.UserID,
        assignment.CardUID,
        assignment.AssignedAt.Format("2006-01-02 15:04:05"),
        assignment.IsActive,
    ).Scan(&assignment.ID)

    if err != nil {
        return fmt.Errorf("error creating assignment: %v", err)
    }
    return nil
}

func (pg *PostgreSQL) GetAssignmentByID(id int) (entities.NfcCardAssignment, error) {
    var assignment entities.NfcCardAssignment
    query := `
        SELECT id, user_id, card_uid, assigned_at, is_active
        FROM nfc_card_assignments
        WHERE id = $1`

    var assignedAtStr string
    err := pg.conn.DB.QueryRow(query, id).Scan(
        &assignment.ID,
        &assignment.UserID,
        &assignment.CardUID,
        &assignedAtStr,
        &assignment.IsActive,
    )

    if err == sql.ErrNoRows {
        return assignment, fmt.Errorf("assignment not found")
    }
    if err != nil {
        return assignment, fmt.Errorf("error getting assignment: %v", err)
    }

    assignment.AssignedAt, _ = time.Parse("2006-01-02 15:04:05", assignedAtStr)
    return assignment, nil
}

func (pg *PostgreSQL) GetAssignmentsByUserID(userID int) ([]entities.NfcCardAssignment, error) {
    query := `
        SELECT id, user_id, card_uid, assigned_at, is_active
        FROM nfc_card_assignments
        WHERE user_id = $1`

    return pg.queryAssignments(query, userID)
}

func (pg *PostgreSQL) GetAssignmentByCardUID(cardUID string) (entities.NfcCardAssignment, error) {
    var assignment entities.NfcCardAssignment
    query := `
        SELECT id, user_id, card_uid, assigned_at, is_active
        FROM nfc_card_assignments
        WHERE card_uid = $1 AND is_active = true`

    var assignedAtStr string
    err := pg.conn.DB.QueryRow(query, cardUID).Scan(
        &assignment.ID,
        &assignment.UserID,
        &assignment.CardUID,
        &assignedAtStr,
        &assignment.IsActive,
    )

    if err == sql.ErrNoRows {
        return assignment, fmt.Errorf("active assignment not found for card")
    }
    if err != nil {
        return assignment, fmt.Errorf("error getting assignment: %v", err)
    }

    assignment.AssignedAt, _ = time.Parse("2006-01-02 15:04:05", assignedAtStr)
    return assignment, nil
}

func (pg *PostgreSQL) UpdateAssignment(assignment *entities.NfcCardAssignment) error {
    query := `
        UPDATE nfc_card_assignments
        SET user_id = $1, card_uid = $2, is_active = $3
        WHERE id = $4`

    result, err := pg.conn.ExecutePreparedQuery(
        query,
        assignment.UserID,
        assignment.CardUID,
        assignment.IsActive,
        assignment.ID,
    )
    if err != nil {
        return fmt.Errorf("error updating assignment: %v", err)
    }
    if result == nil {
        return fmt.Errorf("assignment not found")
    }
    return nil
}

func (pg *PostgreSQL) DeactivateAssignment(id int) error {
    query := `
        UPDATE nfc_card_assignments
        SET is_active = false
        WHERE id = $1`

    result, err := pg.conn.ExecutePreparedQuery(query, id)
    if err != nil {
        return fmt.Errorf("error deactivating assignment: %v", err)
    }
    if result == nil {
        return fmt.Errorf("assignment not found")
    }
    return nil
}

func (pg *PostgreSQL) GetAllAssignments() ([]entities.NfcCardAssignment, error) {
    query := `
        SELECT id, user_id, card_uid, assigned_at, is_active
        FROM nfc_card_assignments`

    return pg.queryAssignments(query)
}

func (pg *PostgreSQL) queryAssignments(query string, args ...interface{}) ([]entities.NfcCardAssignment, error) {
    rows, err := pg.conn.DB.Query(query, args...)
    if err != nil {
        return nil, fmt.Errorf("error querying assignments: %v", err)
    }
    defer rows.Close()

    var assignments []entities.NfcCardAssignment
    for rows.Next() {
        var assignment entities.NfcCardAssignment
        var assignedAtStr string

        err := rows.Scan(
            &assignment.ID,
            &assignment.UserID,
            &assignment.CardUID,
            &assignedAtStr,
            &assignment.IsActive,
        )
        if err != nil {
            return nil, fmt.Errorf("error scanning assignment: %v", err)
        }

        assignment.AssignedAt, _ = time.Parse("2006-01-02 15:04:05", assignedAtStr)
        assignments = append(assignments, assignment)
    }

    return assignments, nil
}