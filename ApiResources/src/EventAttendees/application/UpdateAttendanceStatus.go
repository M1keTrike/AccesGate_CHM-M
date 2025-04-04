package application

import (
    "api_resources/src/EventAttendees/domain"
)

type UpdateAttendanceStatus struct {
    repo domain.EventAttendeeRepository
}

func NewUpdateAttendanceStatus(repo domain.EventAttendeeRepository) *UpdateAttendanceStatus {
    return &UpdateAttendanceStatus{repo: repo}
}

func (uc *UpdateAttendanceStatus) Execute(eventID, userID int, attended bool) error {
    return uc.repo.UpdateAttendanceStatus(eventID, userID, attended)
}