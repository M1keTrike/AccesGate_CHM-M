package application

import (
    "api_resources/src/NfcCardAssignments/domain"
    "api_resources/src/NfcCardAssignments/domain/entities"
)

type GetAllAssignments struct {
    repo domain.NfcCardAssignmentRepository
}

func NewGetAllAssignments(repo domain.NfcCardAssignmentRepository) *GetAllAssignments {
    return &GetAllAssignments{repo: repo}
}

func (uc *GetAllAssignments) Execute() ([]entities.NfcCardAssignment, error) {
    return uc.repo.GetAllAssignments()
}