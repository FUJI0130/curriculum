package rdbimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/jmoiron/sqlx"
)

type tagRepositoryImpl struct {
	Conn *sqlx.DB
}

type tagRequest struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewTagRepository(conn *sqlx.DB) tagdm.TagRepository {
	return &tagRepositoryImpl{Conn: conn}
}

func (repo *tagRepositoryImpl) Store(ctx context.Context, tag *tagdm.Tag) error {

	query := "INSERT INTO tags (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)"
	_, err := repo.Conn.Exec(query, tag.ID(), tag.Name(), tag.CreatedAt().DateTime(), tag.UpdatedAt().DateTime())
	if err != nil {

		return err
	}

	return nil
}

func (repo *tagRepositoryImpl) FindByName(ctx context.Context, name string) (*tagdm.Tag, error) {
	query := "SELECT * FROM tags WHERE name = ?"
	var tempTag tagRequest
	err := repo.Conn.Get(&tempTag, query, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, tagdm.ErrTagNotFound
		}
		return nil, fmt.Errorf("tag_repository_impl FindByName database error: %v", err)
	}

	tagID := tagdm.TagID(tempTag.ID)
	tag, err := tagdm.ReconstructTag(tagID, tempTag.Name, tempTag.CreatedAt, tempTag.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("error converting tagRequest to Tag: %v", err)
	}

	return tag, nil
}

func (repo *tagRepositoryImpl) FindByNames(ctx context.Context, names []string) ([]*tagdm.Tag, error) {
	query := "SELECT * FROM tags WHERE name IN (?)"
	var tempTags []tagRequest
	query, args, err := sqlx.In(query, names)
	if err != nil {
		return nil, err
	}

	err = repo.Conn.Select(&tempTags, query, args...)
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	var tags []*tagdm.Tag
	for _, tempTag := range tempTags {
		tagID := tagdm.TagID(tempTag.ID)
		tag, err := tagdm.ReconstructTag(tagID, tempTag.Name, tempTag.CreatedAt, tempTag.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error converting tagRequest to Tag: %v", err)
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (repo *tagRepositoryImpl) FindByID(ctx context.Context, id string) (*tagdm.Tag, error) {
	query := "SELECT * FROM tags WHERE id = ?"
	var tempTag tagRequest
	err := repo.Conn.Get(&tempTag, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, tagdm.ErrTagNotFound
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	tagID := tagdm.TagID(tempTag.ID)
	tag, err := tagdm.ReconstructTag(tagID, tempTag.Name, tempTag.CreatedAt, tempTag.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error converting tagRequest to Tag: %v", err)
	}

	return tag, nil
}
