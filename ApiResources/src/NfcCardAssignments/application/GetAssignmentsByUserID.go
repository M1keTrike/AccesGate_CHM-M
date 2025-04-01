package application

import (
    "api_resources/src/NfcCardAssignments/domain"
    "api_resources/src/NfcCardAssignments/domain/entities"
)

type GetAssignmentsByUserID struct {
    repo domain.NfcCardAssignmentRepository
}

func NewGetAssignmentsByUserID(repo domain.NfcCardAssignmentRepository) *GetAssignmentsByUserID {
    return &GetAssignmentsByUserID{repo: repo}
}

func (uc *GetAssignmentsByUserID) Execute(userID int) ([]entities.NfcCardAssignment, error) {
    return uc.repo.GetAssignmentsByUserID(userID)
}