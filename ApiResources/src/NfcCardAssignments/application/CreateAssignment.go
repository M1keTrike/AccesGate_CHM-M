package application

import (
    "api_resources/src/NfcCardAssignments/domain"
    "api_resources/src/NfcCardAssignments/domain/entities"
)

type CreateAssignment struct {
    repo domain.NfcCardAssignmentRepository
}

func NewCreateAssignment(repo domain.NfcCardAssignmentRepository) *CreateAssignment {
    return &CreateAssignment{repo: repo}
}

func (uc *CreateAssignment) Execute(assignment *entities.NfcCardAssignment) error {
    return uc.repo.CreateAssignment(assignment)
}