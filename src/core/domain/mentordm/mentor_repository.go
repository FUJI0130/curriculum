// src/core/domain/mentordm/mentor_repository.go

package mentordm

type MentorRepository interface {
	FindAll() ([]*Mentor, error)
}
