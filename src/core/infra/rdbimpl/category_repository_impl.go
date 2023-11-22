package rdbimpl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
	"github.com/FUJI0130/curriculum/src/core/infra/datamodel"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type categoryRepositoryImpl struct{}

func NewCategoryRepository() categorydm.CategoryRepository {
	return &categoryRepositoryImpl{}
}

func (repo *categoryRepositoryImpl) Store(ctx context.Context, category *categorydm.Category) error {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return errors.New("no transaction found in context")
	}
	query := "INSERT INTO categories (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)"
	_, err := conn.Exec(query, category.ID(), category.Name(), category.CreatedAt().DateTime(), category.UpdatedAt().DateTime())

	if err != nil {
		return customerrors.WrapInternalServerError(err, "カテゴリの保存に失敗しました")
	}

	return nil
}

func (repo *categoryRepositoryImpl) FindByName(ctx context.Context, name string) (*categorydm.Category, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return nil, errors.New("コンテキスト内でトランザクションが見つかりません")
	}
	query := "SELECT * FROM categories WHERE name = ?"
	var tempCategory datamodel.Category
	err := conn.Get(&tempCategory, query, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// カテゴリが見つからなかった場合、nil と sql.ErrNoRows を返す
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	categoryID, err := categorydm.NewCategoryIDFromString(tempCategory.ID)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "IDのパースに失敗しました")

	}

	category, err := categorydm.GenWhenFetch(categoryID, tempCategory.Name, tempCategory.CreatedAt)

	if err != nil {
		return nil, customerrors.WrapInternalServerError(err, "カテゴリの再構築に失敗しました")
	}

	return category, nil
}

func (repo *categoryRepositoryImpl) FindByID(ctx context.Context, id string) (*categorydm.Category, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return nil, errors.New("コンテキスト内でトランザクションが見つかりません")
	}
	query := "SELECT * FROM categories WHERE id = ?"
	var tempCategory datamodel.Category
	err := conn.Get(&tempCategory, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.WrapNotFoundError(err, "カテゴリが見つかりません")
		}
		return nil, err
	}

	categoryID, err := categorydm.NewCategoryIDFromString(tempCategory.ID)
	if err != nil {
		return nil, customerrors.WrapUnprocessableEntityError(err, "IDのパースに失敗しました")
	}

	category, err := categorydm.GenWhenFetch(categoryID, tempCategory.Name, tempCategory.CreatedAt)

	if err != nil {
		return nil, customerrors.WrapInternalServerError(err, "カテゴリの再構築に失敗しました")
	}

	return category, nil
}
