package application

import (
	"api_resources/src/Devices/domain"
	"api_resources/src/Devices/domain/entities"
	"time"
)

type UpdateDevice struct {
	repo domain.DeviceRepository
}

func NewUpdateDevice(repo domain.DeviceRepository) *UpdateDevice {
	return &UpdateDevice{repo: repo}
}

func (uc *UpdateDevice) Execute(device *entities.Device) error {
	device.UpdatedAt = time.Now()
	return uc.repo.UpdateDevice(device)
}
