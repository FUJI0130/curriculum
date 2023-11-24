package rdbimpl

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/infra/datamodel"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type userRepositoryImpl struct{}

func NewUserRepository() userdm.UserRepository {
	return &userRepositoryImpl{}
}

func (repo *userRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {

	conn, exists := ctx.Value("Conn").(dbOperator)

	if !exists {
		return errors.New("no transaction found in context")
	}

	queryUser := "INSERT INTO users (id, name, email, password, profile, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := conn.Exec(queryUser, user.ID().String(), user.Name(), user.Email(), user.Password(), user.Profile(), user.CreatedAt().DateTime(), user.UpdatedAt().DateTime())
	if err != nil {
		return customerrors.WrapInternalServerError(err, "Failed to store user")
	}

	for _, skill := range user.Skills() {
		querySkill := "INSERT INTO skills (id,tag_id,user_id,created_at,updated_at, evaluation, years) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err = conn.Exec(querySkill, skill.ID().String(), skill.TagID().String(), user.ID().String(), skill.CreatedAt().DateTime(), skill.UpdatedAt().DateTime(), skill.Evaluation().Value(), skill.Year().Value())
		if err != nil {
			return customerrors.WrapInternalServerError(err, "Failed to store skill")
		}
	}

	for _, career := range user.Careers() {
		queryCareer := "INSERT INTO careers (id,user_id, detail, ad_from, ad_to, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err = conn.Exec(queryCareer, career.ID().String(), user.ID().String(), career.Detail(), career.AdFrom(), career.AdTo(), career.CreatedAt().DateTime(), career.UpdatedAt().DateTime())
		if err != nil {
			log.Printf("Failed to store career")
			return customerrors.WrapInternalServerError(err, "Failed to store career")
		}
	}
	return nil
}

func (repo *userRepositoryImpl) FindByUserName(ctx context.Context, name string) (*userdm.User, error) {
	log.Printf("FindByUserName: name: %s", name)
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		log.Printf("no transaction found in context")
		return nil, errors.New("no transaction found in context")
	}

	// ユーザー情報の取得
	userQuery := "SELECT * FROM users WHERE name = ?"
	var tempUser datamodel.User
	if err := conn.Get(&tempUser, userQuery, name); err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.WrapNotFoundError(err, "User not found by name")
		}
		return nil, err
	}

	// スキル情報の取得
	skillQuery := "SELECT * FROM skills WHERE user_id = ?"
	var tempSkills []datamodel.Skill
	if err := conn.Select(&tempSkills, skillQuery, tempUser.ID); err != nil {
		return nil, err
	}
	var skills []userdm.Skill
	for _, tempSkill := range tempSkills {
		skill, err := userdm.ReconstructSkill(tempSkill.ID, tempSkill.TagID, tempSkill.Evaluation, tempSkill.Years, tempSkill.CreatedAt, tempSkill.UpdatedAt)
		if err != nil {
			return nil, err
		}
		skills = append(skills, *skill)
	}

	// キャリア情報の取得
	careerQuery := "SELECT * FROM careers WHERE user_id = ?"
	var tempCareers []datamodel.Career
	if err := conn.Select(&tempCareers, careerQuery, tempUser.ID); err != nil {
		return nil, err
	}
	var careers []userdm.Career
	for _, tempCareer := range tempCareers {
		career, err := userdm.ReconstructCareer(tempCareer.ID, tempCareer.Detail, tempCareer.AdFrom, tempCareer.AdTo, tempCareer.CreatedAt, tempCareer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		careers = append(careers, *career)
	}

	// 完全なユーザー集約の再構築
	return userdm.ReconstructEntity(tempUser.ID, tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, skills, careers, tempUser.CreatedAt)
}

func (repo *userRepositoryImpl) FindByUserID(ctx context.Context, userID string) (*userdm.User, error) {
	log.Printf("findUserByUserID: userID: %s", userID)
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		log.Printf("no transaction found in context")
		return nil, errors.New("no transaction found in context")
	}

	// ユーザー情報の取得
	query := "SELECT * FROM users WHERE id = ?"
	var tempUser datamodel.User
	if err := conn.Get(&tempUser, query, userID); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("sql.ErrNoRows")
			return nil, customerrors.WrapNotFoundError(err, "User not found by userID")
		}
		return nil, err
	}

	// スキル情報の取得
	skillQuery := "SELECT * FROM skills WHERE user_id = ?"
	var tempSkills []datamodel.Skill
	if err := conn.Select(&tempSkills, skillQuery, userID); err != nil {
		return nil, err
	}
	var skills []userdm.Skill
	for _, tempSkill := range tempSkills {
		skill, err := userdm.ReconstructSkill(tempSkill.ID, tempSkill.TagID, tempSkill.Evaluation, tempSkill.Years, tempSkill.CreatedAt, tempSkill.UpdatedAt)
		if err != nil {
			return nil, err
		}
		skills = append(skills, *skill)
	}

	// キャリア情報の取得
	careerQuery := "SELECT * FROM careers WHERE user_id = ?"
	var tempCareers []datamodel.Career
	if err := conn.Select(&tempCareers, careerQuery, userID); err != nil {
		return nil, err
	}
	var careers []userdm.Career
	for _, tempCareer := range tempCareers {
		career, err := userdm.ReconstructCareer(tempCareer.ID, tempCareer.Detail, tempCareer.AdFrom, tempCareer.AdTo, tempCareer.CreatedAt, tempCareer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		careers = append(careers, *career)
	}

	// 完全なユーザー集約の再構築
	return userdm.ReconstructEntity(tempUser.ID, tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, skills, careers, tempUser.CreatedAt)
}

func (repo *userRepositoryImpl) Update(ctx context.Context, user *userdm.User) error {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return errors.New("no transaction found in context")
	}

	queryUser := "UPDATE users SET name = ?, email = ?, password = ?, profile = ?, updated_at = ? WHERE id = ?"
	if _, err := conn.Exec(queryUser, user.Name(), user.Email(), user.Password(), user.Profile(), time.Now(), user.ID().String()); err != nil {
		return err
	}

	for _, skill := range user.Skills() {
		querySkill := "UPDATE skills SET tag_id = ?, evaluation = ?, years = ?, updated_at = ? WHERE id = ?"
		if _, err := conn.Exec(querySkill, skill.TagID().String(), skill.Evaluation().Value(), skill.Year().Value(), time.Now(), skill.ID().String()); err != nil {
			return err
		}
	}

	for _, career := range user.Careers() {
		queryCareer := "UPDATE careers SET detail = ?, ad_from = ?, ad_to = ?, updated_at = ? WHERE id = ?"
		if _, err := conn.Exec(queryCareer, career.Detail(), career.AdFrom(), career.AdTo(), time.Now(), career.ID().String()); err != nil {
			return err
		}
	}

	return nil
}
