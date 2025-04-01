package application

import "api_resources/src/Devices/domain"

type DeleteDevice struct {
	repo domain.DeviceRepository
}

func NewDeleteDevice(repo domain.DeviceRepository) *DeleteDevice {
	return &DeleteDevice{repo: repo}
}

func (uc *DeleteDevice) Execute(id int) error {
	return uc.repo.DeleteDevice(id)
}
