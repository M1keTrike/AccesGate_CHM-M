package infraestructure

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    "api_resources/src/core"
    "api_resources/src/Users/domain/entities"
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

func (pg *PostgreSQL) GetUserByID(userID int) (entities.User, error) {
    user := entities.User{}
    query := "SELECT id, name, email, password_hash, role, created_at FROM users WHERE id = $1"
    row := pg.conn.DB.QueryRow(query, userID)

    var createdAtStr string
    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &createdAtStr)
    if err != nil {
        if err == sql.ErrNoRows {
            return user, fmt.Errorf("usuario con id %d no encontrado", userID)
        }
        return user, fmt.Errorf("error al obtener usuario: %v", err)
    }
    user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
    if err != nil {
        return user, fmt.Errorf("error al parsear la fecha de creación: %v", err)
    }

    return user, nil
}

func (pg *PostgreSQL) CreateUser(user *entities.User) error {
    query := "INSERT INTO users (name, email, password_hash, role, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
    err := pg.conn.DB.QueryRow(query, user.Name, user.Email, user.PasswordHash, user.Role, user.CreatedAt).Scan(&user.ID)
    if err != nil {
        return fmt.Errorf("error al crear usuario: %v", err)
    }
    return nil
}

func (pg *PostgreSQL) UpdateUser(user *entities.User) error {
    query := "UPDATE users SET name = $1, email = $2, password_hash = $3, role = $4 WHERE id = $5"
    _, err := pg.conn.ExecutePreparedQuery(query, user.Name, user.Email, user.PasswordHash, user.Role, user.ID)
    if err != nil {
        return fmt.Errorf("error al actualizar usuario: %v", err)
    }
    return nil
}

func (pg *PostgreSQL) DeleteUser(userID int) error {
    query := "DELETE FROM users WHERE id = $1"
    _, err := pg.conn.ExecutePreparedQuery(query, userID)
    if err != nil {
        return fmt.Errorf("error al eliminar usuario: %v", err)
    }
    return nil
}

func (pg *PostgreSQL) GetAllUsers() ([]entities.User, error) {
    users := []entities.User{}
    query := "SELECT id, name, email, password_hash, role, created_at FROM users"
    rows, err := pg.conn.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error al obtener usuarios: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var createdAtStr string
        user := entities.User{}
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &createdAtStr); err != nil {
            return nil, fmt.Errorf("error al escanear usuarios: %v", err)
        }
        user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
        if err != nil {
            return nil, fmt.Errorf("error al parsear la fecha de creación: %v", err)
        }
        users = append(users, user)
    }

    return users, nil
}

func (pg *PostgreSQL) GetUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	query := "SELECT id, name, email, password_hash, role, created_at FROM users WHERE email = $1"
	row := pg.conn.DB.QueryRow(query, email)

	var createdAtStr string
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &createdAtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con email %s no encontrado", email)
		}
		return nil, fmt.Errorf("error al obtener usuario: %v", err)
	}
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la fecha de creación: %v", err)
	}

	return user, nil
}