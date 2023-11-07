// src/infra/rdbimpl/mentor_repository_impl.go

package rdbimpl

import (
	"github.com/FUJI0130/curriculum/src/core/domain/mentordm"
)

type MentorRepositoryImpl struct {
	SqlHandler
}

func NewMentorRepository(sqlHandler SqlHandler) mentordm.MentorRepository {
	return &MentorRepositoryImpl{sqlHandler}
}

func (repo *MentorRepositoryImpl) FindAll() ([]*mentordm.Mentor, error) {
	const query = `
        SELECT id, name, profile, created_at, updated_at
        FROM users
        WHERE /* 何らかの条件でメンターをフィルターする */
    `
	rows, err := repo.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mentors []*mentordm.Mentor
	for rows.Next() {
		var m mentordm.Mentor
		err := rows.Scan(&m.ID, &m.Name, &m.Profile, &m.CreatedAt, &m.UpdatedAt)
		if err != nil {
			return nil, err
		}
		mentors = append(mentors, &m)
	}

	return mentors, nil
}
