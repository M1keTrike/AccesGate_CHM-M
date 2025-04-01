package application

import (
    "api_resources/src/NfcCardAssignments/domain"
    "api_resources/src/NfcCardAssignments/domain/entities"
)

type GetAssignmentByID struct {
    repo domain.NfcCardAssignmentRepository
}

func NewGetAssignmentByID(repo domain.NfcCardAssignmentRepository) *GetAssignmentByID {
    return &GetAssignmentByID{repo: repo}
}

func (uc *GetAssignmentByID) Execute(id int) (entities.NfcCardAssignment, error) {
    return uc.repo.GetAssignmentByID(id)
}