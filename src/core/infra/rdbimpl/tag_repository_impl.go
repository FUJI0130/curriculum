package rdbimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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
	// CreatedAt string `db:"created_at"`
	// UpdatedAt string `db:"updated_at"`
}

func NewTagRepository(conn *sqlx.DB) tagdm.TagRepository {
	return &tagRepositoryImpl{Conn: conn}
}

func (repo *tagRepositoryImpl) Store(ctx context.Context, tag *tagdm.Tag) error {

	log.Printf("[DEBUG] tag_repository_impl Store")
	query := "INSERT INTO tags (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)"
	_, err := repo.Conn.Exec(query, tag.ID(), tag.Name(), tag.CreatedAt().DateTime(), tag.UpdatedAt().DateTime())
	if err != nil {

		log.Printf("[DEBUG] into error handling tag_repository_impl Store err is : %s", err)
		log.Printf("[DEBUG] tag_repository_impl Store")
		return err
	}

	log.Printf("[DEBUG] after error handling tag_repository_impl Store")
	return nil
}

func (repo *tagRepositoryImpl) FindByName(ctx context.Context, name string) (*tagdm.Tag, error) {

	log.Printf("[DEBUG] before tag_repository_impl FindByName")
	query := "SELECT * FROM tags WHERE name = ?"
	var tempTag tagRequest
	log.Printf("[DEBUG] before repo.Conn.Get(&tempTag, query, name)")
	err := repo.Conn.Get(&tempTag, query, name)
	if err != nil {

		log.Printf("[DEBUG] into error handling repo.Conn.Get 1")
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("[DEBUG] into error handling repo.Conn.Get 2 ErrTagNotFound")
			return nil, tagdm.ErrTagNotFound
		}
		log.Printf("[DEBUG] Not ErrNoRows")
		return nil, fmt.Errorf("database error: %v", err)
	}

	log.Printf("[DEBUG] after select query")
	tagID := tagdm.TagID(tempTag.ID)
	tag, err := tagdm.ReconstructTag(tagID, tempTag.Name, tempTag.CreatedAt, tempTag.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error converting tagRequest to Tag: %v", err)
	}

	return tag, nil
}

func (repo *tagRepositoryImpl) FindByID(ctx context.Context, id string) (*tagdm.Tag, error) {
	query := "SELECT * FROM tags WHERE id = ?"
	var tempTag tagRequest
	err := repo.Conn.Get(&tempTag, query, id)
	if err != nil {
		log.Println("[DEBUG] FindByID error:", err)
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Error in FindByID:", err)
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

func (repo *tagRepositoryImpl) CreateNewTag(ctx context.Context, name string) (*tagdm.Tag, error) {
	// 新規タグを作成

	log.Printf("[DEBUG] before CreateNewTag")
	newTag, err := tagdm.NewTag(name, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] before CreateNewTag")
	// データベースに新規タグを保存
	err = repo.Store(ctx, newTag)
	if err != nil {
		return nil, err
	}

	return newTag, nil
}
