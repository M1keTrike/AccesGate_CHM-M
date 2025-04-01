package domain

import "api_resources/src/Devices/domain/entities"

type DeviceRepository interface {
	CreateDevice(device *entities.Device) error
	GetAllDevices() ([]entities.Device, error)
	GetDeviceByID(id int) (*entities.Device, error)
	GetDeviceByHardwareID(hardwareID string) (*entities.Device, error)
	UpdateDevice(device *entities.Device) error
	DeleteDevice(id int) error
}
