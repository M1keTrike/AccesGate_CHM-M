package application

import (
	"api_resources/src/Devices/domain"
	"api_resources/src/Devices/domain/entities"
)

type GetDeviceByHardwareID struct {
	repo domain.DeviceRepository
}

func NewGetDeviceByHardwareID(repo domain.DeviceRepository) *GetDeviceByHardwareID {
	return &GetDeviceByHardwareID{repo: repo}
}

func (uc *GetDeviceByHardwareID) Execute(hardwareID string) (*entities.Device, error) {
	return uc.repo.GetDeviceByHardwareID(hardwareID)
}
