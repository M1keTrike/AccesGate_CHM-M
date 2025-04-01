package application

import (
	"api_resources/src/Devices/domain"
	"api_resources/src/Devices/domain/entities"
	"time"
)

type CreateDevice struct {
	repo domain.DeviceRepository
}

func NewCreateDevice(repo domain.DeviceRepository) *CreateDevice {
	return &CreateDevice{repo: repo}
}

func (uc *CreateDevice) Execute(device *entities.Device) error {
	now := time.Now()
	device.RegisteredAt = now
	device.UpdatedAt = now
	return uc.repo.CreateDevice(device)
}
