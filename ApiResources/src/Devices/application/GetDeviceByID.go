package application

import (
	"api_resources/src/Devices/domain"
	"api_resources/src/Devices/domain/entities"
)

type GetDeviceByID struct {
	repo domain.DeviceRepository
}

func NewGetDeviceByID(repo domain.DeviceRepository) *GetDeviceByID {
	return &GetDeviceByID{repo: repo}
}

func (uc *GetDeviceByID) Execute(id int) (*entities.Device, error) {
	return uc.repo.GetDeviceByID(id)
}
