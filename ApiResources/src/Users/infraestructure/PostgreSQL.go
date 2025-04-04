package infraestructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"api_resources/src/Users/domain/entities"
	"api_resources/src/core"
)

type PostgreSQL struct {
	conn *core.Conn_PostgreSQL
}

func NewPostgreSQL() *PostgreSQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("[DB INIT] Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &PostgreSQL{conn: conn}
}

func (pg *PostgreSQL) GetUserByID(userID int) (entities.User, error) {
	user := entities.User{}
	query := `
		SELECT id, name, email, password_hash, role,
		       created_at, updated_at, fingerprint_id, biometric_auth, created_by
		FROM users
		WHERE id = $1`
	row := pg.conn.DB.QueryRow(query, userID)

	var createdAt, updatedAt sql.NullTime
	var fingerprintID sql.NullInt16

	err := row.Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role,
		&createdAt, &updatedAt, &fingerprintID, &user.BiometricAuth, &user.CreatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[GetUserByID] Usuario con ID %d no encontrado", userID)
			return user, fmt.Errorf("usuario no encontrado")
		}
		log.Printf("[GetUserByID] Error al obtener usuario con ID %d: %v", userID, err)
		return user, fmt.Errorf("error al obtener usuario")
	}

	if createdAt.Valid {
		user.CreatedAt = createdAt.Time
	}
	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.Time
	}
	if fingerprintID.Valid {
		user.FingerprintID = fingerprintID.Int16
	}

	log.Printf("[GetUserByID] Usuario con ID %d obtenido correctamente", userID)
	return user, nil
}

func (pg *PostgreSQL) CreateUser(user *entities.User) error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	query := `
		INSERT INTO users (
			name, email, password_hash, role,
			created_at, updated_at, fingerprint_id, biometric_auth, created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`
	err := pg.conn.DB.QueryRow(
		query,
		user.Name, user.Email, user.PasswordHash, user.Role,
		user.CreatedAt, user.UpdatedAt, user.FingerprintID, user.BiometricAuth, user.CreatedBy,
	).Scan(&user.ID)
	if err != nil {
		log.Printf("[CreateUser] Error al crear usuario '%s': %v", user.Email, err)
		return err
	}

	log.Printf("[CreateUser] Usuario '%s' creado con ID %d", user.Email, user.ID)
	return nil
}

func (pg *PostgreSQL) UpdateUser(user *entities.User) error {
	user.UpdatedAt = time.Now()

	query := `
		UPDATE users SET
			name = $1,
			email = $2,
			password_hash = $3,
			role = $4,
			updated_at = $5,
			fingerprint_id = $6,
			biometric_auth = $7,
			created_by = $8
		WHERE id = $9`

	_, err := pg.conn.ExecutePreparedQuery(
		query,
		user.Name, user.Email, user.PasswordHash, user.Role,
		user.UpdatedAt, user.FingerprintID, user.BiometricAuth, user.CreatedBy, user.ID,
	)
	if err != nil {
		log.Printf("[UpdateUser] Error al actualizar usuario con ID %d: %v", user.ID, err)
		return fmt.Errorf("error al actualizar usuario")
	}

	log.Printf("[UpdateUser] Usuario con ID %d actualizado", user.ID)
	return nil
}

func (pg *PostgreSQL) DeleteUser(userID int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := pg.conn.ExecutePreparedQuery(query, userID)
	if err != nil {
		log.Printf("[DeleteUser] Error al eliminar usuario con ID %d: %v", userID, err)
		return fmt.Errorf("error al eliminar usuario")
	}

	log.Printf("[DeleteUser] Usuario con ID %d eliminado", userID)
	return nil
}

func (pg *PostgreSQL) GetAllUsers() ([]entities.User, error) {
	users := []entities.User{}
	query := `
		SELECT id, name, email, password_hash, role,
		       created_at, updated_at, fingerprint_id, biometric_auth, created_by
		FROM users`
	rows, err := pg.conn.DB.Query(query)
	if err != nil {
		log.Printf("[GetAllUsers] Error al ejecutar consulta: %v", err)
		return nil, fmt.Errorf("error al obtener usuarios")
	}
	defer rows.Close()

	for rows.Next() {
		user := entities.User{}
		var createdAt, updatedAt sql.NullTime
		var fingerprintID sql.NullInt16

		err := rows.Scan(
			&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role,
			&createdAt, &updatedAt, &fingerprintID, &user.BiometricAuth, &user.CreatedBy,
		)
		if err != nil {
			log.Printf("[GetAllUsers] Error al escanear fila: %v", err)
			return nil, fmt.Errorf("error al escanear usuarios")
		}

		if createdAt.Valid {
			user.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.Time
		}
		if fingerprintID.Valid {
			user.FingerprintID = fingerprintID.Int16
		}

		users = append(users, user)
	}

	log.Printf("[GetAllUsers] Se obtuvieron %d usuarios", len(users))
	return users, nil
}

func (pg *PostgreSQL) GetUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	query := `
		SELECT id, name, email, password_hash, role,
		       created_at, updated_at, fingerprint_id, biometric_auth
		FROM users
		WHERE email = $1`
	row := pg.conn.DB.QueryRow(query, email)

	var createdAt, updatedAt sql.NullTime
	var fingerprintID sql.NullInt16

	err := row.Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role,
		&createdAt, &updatedAt, &fingerprintID, &user.BiometricAuth,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[GetUserByEmail] Usuario con email %s no encontrado", email)
			return nil, fmt.Errorf("usuario no encontrado")
		}
		log.Printf("[GetUserByEmail] Error al obtener usuario con email %s: %v", email, err)
		return nil, fmt.Errorf("error al obtener usuario")
	}

	if createdAt.Valid {
		user.CreatedAt = createdAt.Time
	}
	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.Time
	}
	if fingerprintID.Valid {
		user.FingerprintID = fingerprintID.Int16
	}

	log.Printf("[GetUserByEmail] Usuario '%s' obtenido correctamente", email)
	return user, nil
}

func (pg *PostgreSQL) GetUsersByRole(role string) ([]entities.User, error) {
	users := []entities.User{}
	query := `
        SELECT id, name, email, password_hash, role,
               created_at, updated_at, fingerprint_id, biometric_auth, created_by
        FROM users 
        WHERE role = $1`

	rows, err := pg.conn.DB.Query(query, role)
	if err != nil {
		log.Printf("[GetUsersByRole] Error al obtener usuarios con rol %s: %v", role, err)
		return nil, fmt.Errorf("error al obtener usuarios por rol")
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		var createdAt, updatedAt sql.NullTime
		var fingerprintID sql.NullInt16

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&user.Role,
			&createdAt,
			&updatedAt,
			&fingerprintID,
			&user.BiometricAuth,
			&user.CreatedBy,
		)
		if err != nil {
			log.Printf("[GetUsersByRole] Error al escanear usuario: %v", err)
			return nil, fmt.Errorf("error al procesar usuarios")
		}

		if createdAt.Valid {
			user.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.Time
		}
		if fingerprintID.Valid {
			user.FingerprintID = fingerprintID.Int16
		}

		users = append(users, user)
	}

	log.Printf("[GetUsersByRole] Se obtuvieron %d usuarios con rol %s", len(users), role)
	return users, nil
}

func (pg *PostgreSQL) GetUsersByCreatedBy(createdBy int) ([]entities.User, error) {
	users := []entities.User{}
	query := `
        SELECT id, name, email, password_hash, role,
               created_at, updated_at, fingerprint_id, biometric_auth, created_by
        FROM users 
        WHERE created_by = $1`

	rows, err := pg.conn.DB.Query(query, createdBy)
	if err != nil {
		log.Printf("[GetUsersByCreatedBy] Error al obtener usuarios creados por %d: %v", createdBy, err)
		return nil, fmt.Errorf("error al obtener usuarios por creador")
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		var createdAt, updatedAt sql.NullTime
		var fingerprintID sql.NullInt16

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&user.Role,
			&createdAt,
			&updatedAt,
			&fingerprintID,
			&user.BiometricAuth,
			&user.CreatedBy,
		)
		if err != nil {
			log.Printf("[GetUsersByCreatedBy] Error al escanear usuario: %v", err)
			return nil, fmt.Errorf("error al procesar usuarios")
		}

		if createdAt.Valid {
			user.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.Time
		}
		if fingerprintID.Valid {
			user.FingerprintID = fingerprintID.Int16
		}

		users = append(users, user)
	}

	log.Printf("[GetUsersByCreatedBy] Se obtuvieron %d usuarios creados por el usuario %d", len(users), createdBy)
	return users, nil
}
