package application

import (
    "api_resources/src/NfcCardAssignments/domain"
    "api_resources/src/NfcCardAssignments/domain/entities"
)

type UpdateAssignment struct {
    repo domain.NfcCardAssignmentRepository
}

func NewUpdateAssignment(repo domain.NfcCardAssignmentRepository) *UpdateAssignment {
    return &UpdateAssignment{repo: repo}
}

func (uc *UpdateAssignment) Execute(assignment *entities.NfcCardAssignment) error {
    return uc.repo.UpdateAssignment(assignment)
}