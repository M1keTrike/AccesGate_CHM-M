package application

import (
	"api_resources/src/Devices/domain"
	"api_resources/src/Devices/domain/entities"
)

type GetAllDevices struct {
	repo domain.DeviceRepository
}

func NewGetAllDevices(repo domain.DeviceRepository) *GetAllDevices {
	return &GetAllDevices{repo: repo}
}

func (uc *GetAllDevices) Execute() ([]entities.Device, error) {
	return uc.repo.GetAllDevices()
}
