package application

import (
    "api_resources/src/NfcCardAssignments/domain"
    "api_resources/src/NfcCardAssignments/domain/entities"
)

type GetAssignmentByCardUID struct {
    repo domain.NfcCardAssignmentRepository
}

func NewGetAssignmentByCardUID(repo domain.NfcCardAssignmentRepository) *GetAssignmentByCardUID {
    return &GetAssignmentByCardUID{repo: repo}
}

func (uc *GetAssignmentByCardUID) Execute(cardUID string) (entities.NfcCardAssignment, error) {
    return uc.repo.GetAssignmentByCardUID(cardUID)
}