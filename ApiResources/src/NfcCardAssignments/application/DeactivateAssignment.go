package application

import (
    "api_resources/src/NfcCardAssignments/domain"
)

type DeactivateAssignment struct {
    repo domain.NfcCardAssignmentRepository
}

func NewDeactivateAssignment(repo domain.NfcCardAssignmentRepository) *DeactivateAssignment {
    return &DeactivateAssignment{repo: repo}
}

func (uc *DeactivateAssignment) Execute(id int) error {
    return uc.repo.DeactivateAssignment(id)
}