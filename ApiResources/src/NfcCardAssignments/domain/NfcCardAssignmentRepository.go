package domain

import "api_resources/src/NfcCardAssignments/domain/entities"

type NfcCardAssignmentRepository interface {
    CreateAssignment(assignment *entities.NfcCardAssignment) error
    GetAssignmentByID(id int) (entities.NfcCardAssignment, error)
    GetAssignmentsByUserID(userID int) ([]entities.NfcCardAssignment, error)
    GetAssignmentByCardUID(cardUID string) (entities.NfcCardAssignment, error)
    UpdateAssignment(assignment *entities.NfcCardAssignment) error
    DeactivateAssignment(id int) error
    GetAllAssignments() ([]entities.NfcCardAssignment, error)
}