package infraestructure

import (
	"api_resources/src/Devices/domain/entities"
	"api_resources/src/core"
	"database/sql"
	"fmt"
	"log"
)

type PostgreSQL struct {
	conn *core.Conn_PostgreSQL
}

func NewPostgreSQLEvents() *PostgreSQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("[Devices DB INIT] Error en el pool de conexiones: %v", conn.Err)
	}
	return &PostgreSQL{conn: conn}
}

func (pg *PostgreSQL) CreateDevice(device *entities.Device) error {
	query := `
		INSERT INTO devices (hardware_id, type, status, location, registered_at, updated_at, assigned_to)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`
	err := pg.conn.DB.QueryRow(
		query,
		device.HardwareID, device.Type, device.Status, device.Location,
		device.RegisteredAt, device.UpdatedAt, device.AssignedTo,
	).Scan(&device.ID)

	if err != nil {
		log.Printf("[CreateDevice] Error al crear dispositivo: %v", err)
		return fmt.Errorf("error al crear dispositivo")
	}

	log.Printf("[CreateDevice] Dispositivo creado con ID %d", device.ID)
	return nil
}

func (pg *PostgreSQL) GetAllDevices() ([]entities.Device, error) {
	query := `SELECT id, hardware_id, type, status, location, registered_at, updated_at, assigned_to FROM devices`
	rows, err := pg.conn.DB.Query(query)
	if err != nil {
		log.Printf("[GetAllDevices] Error en consulta: %v", err)
		return nil, fmt.Errorf("error al obtener dispositivos")
	}
	defer rows.Close()

	var devices []entities.Device

	for rows.Next() {
		var d entities.Device
		var regAt, updAt sql.NullTime

		err := rows.Scan(
			&d.ID, &d.HardwareID, &d.Type, &d.Status, &d.Location,
			&regAt, &updAt, &d.AssignedTo,
		)
		if err != nil {
			log.Printf("[GetAllDevices] Error escaneando fila: %v", err)
			continue
		}
		if regAt.Valid {
			d.RegisteredAt = regAt.Time
		}
		if updAt.Valid {
			d.UpdatedAt = updAt.Time
		}
		devices = append(devices, d)
	}

	return devices, nil
}

func (pg *PostgreSQL) GetDeviceByID(id int) (*entities.Device, error) {
	query := `SELECT id, hardware_id, type, status, location, registered_at, updated_at, assigned_to FROM devices WHERE id = $1`
	row := pg.conn.DB.QueryRow(query, id)

	var d entities.Device
	var regAt, updAt sql.NullTime

	err := row.Scan(&d.ID, &d.HardwareID, &d.Type, &d.Status, &d.Location, &regAt, &updAt, &d.AssignedTo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("dispositivo no encontrado")
		}
		log.Printf("[GetDeviceByID] Error: %v", err)
		return nil, fmt.Errorf("error al obtener dispositivo")
	}
	if regAt.Valid {
		d.RegisteredAt = regAt.Time
	}
	if updAt.Valid {
		d.UpdatedAt = updAt.Time
	}
	return &d, nil
}

func (pg *PostgreSQL) GetDeviceByHardwareID(hardwareID string) (*entities.Device, error) {
	query := `SELECT id, hardware_id, type, status, location, registered_at, updated_at, assigned_to FROM devices WHERE hardware_id = $1`
	row := pg.conn.DB.QueryRow(query, hardwareID)

	var d entities.Device
	var regAt, updAt sql.NullTime

	err := row.Scan(&d.ID, &d.HardwareID, &d.Type, &d.Status, &d.Location, &regAt, &updAt, &d.AssignedTo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("dispositivo no encontrado")
		}
		log.Printf("[GetDeviceByHardwareID] Error: %v", err)
		return nil, fmt.Errorf("error al obtener dispositivo")
	}
	if regAt.Valid {
		d.RegisteredAt = regAt.Time
	}
	if updAt.Valid {
		d.UpdatedAt = updAt.Time
	}
	return &d, nil
}

func (pg *PostgreSQL) UpdateDevice(device *entities.Device) error {
	query := `
		UPDATE devices SET hardware_id = $1, type = $2, status = $3, location = $4,
		updated_at = $5, assigned_to = $6 WHERE id = $7
	`

	_, err := pg.conn.ExecutePreparedQuery(
		query,
		device.HardwareID, device.Type, device.Status, device.Location,
		device.UpdatedAt, device.AssignedTo, device.ID,
	)

	if err != nil {
		log.Printf("[UpdateDevice] Error al actualizar dispositivo %d: %v", device.ID, err)
		return fmt.Errorf("error al actualizar dispositivo")
	}

	log.Printf("[UpdateDevice] Dispositivo con ID %d actualizado", device.ID)
	return nil
}

func (pg *PostgreSQL) DeleteDevice(id int) error {
	query := `DELETE FROM devices WHERE id = $1`
	_, err := pg.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("[DeleteDevice] Error al eliminar dispositivo %d: %v", id, err)
		return fmt.Errorf("error al eliminar dispositivo")
	}

	log.Printf("[DeleteDevice] Dispositivo con ID %d eliminado", id)
	return nil
}
