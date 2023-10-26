package rdbimpl

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/jmoiron/sqlx"
)

type userRepositoryImpl struct{}

type userRequest struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Profile   string    `db:"profile"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type skillRequest struct {
	ID         string    `db:"id"`
	TagID      string    `db:"tag_id"`
	UserID     string    `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Evaluation uint8     `db:"evaluation"`
	Years      uint8     `db:"years"`
}

func NewUserRepository() userdm.UserRepository {
	return &userRepositoryImpl{}
}
func (repo *userRepositoryImpl) Store(ctx context.Context, userdomain *userdm.UserDomain) error {

	conn, exists := ctx.Value("Conn").(dbOperator)

	if !exists {
		return errors.New("no transaction found in context")
	}

	queryUser := "INSERT INTO users (id, name, email, password, profile, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := conn.Exec(queryUser, userdomain.User.ID().String(), userdomain.User.Name(), userdomain.User.Email(), userdomain.User.Password(), userdomain.User.Profile(), userdomain.User.CreatedAt().DateTime(), userdomain.User.UpdatedAt().DateTime())
	if err != nil {
		return customerrors.WrapInternalServerError(err, "Failed to store user")
	}

	for _, skill := range userdomain.Skills {
		querySkill := "INSERT INTO skills (id,tag_id,user_id,created_at,updated_at, evaluation, years) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err = conn.Exec(querySkill, skill.ID().String(), skill.TagID().String(), userdomain.User.ID().String(), skill.CreatedAt().DateTime(), skill.UpdatedAt().DateTime(), skill.Evaluation().Value(), skill.Year().Value())
		if err != nil {
			return customerrors.WrapInternalServerError(err, "Failed to store skill")
		}
	}

	for _, career := range userdomain.Careers {
		queryCareer := "INSERT INTO careers (id,user_id, detail, ad_from, ad_to, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err = conn.Exec(queryCareer, career.ID().String(), career.UserID().String(), career.Detail(), career.AdFrom(), career.AdTo(), career.CreatedAt().DateTime(), career.UpdatedAt().DateTime())
		if err != nil {
			log.Printf("Failed to store career")
			return customerrors.WrapInternalServerError(err, "Failed to store career")
		}
	}
	return nil
}

func (repo *userRepositoryImpl) FindByUserName(ctx context.Context, name string) (*userdm.User, error) {

	conn, exists := ctx.Value("Conn").(dbOperator)

	if !exists {
		return nil, errors.New("no transaction found in context")
	}
	query := "SELECT * FROM users WHERE name = ?"

	var tempUser userRequest
	err := conn.Get(&tempUser, query, name)
	if err != nil {
		log.Printf("user Repository FindByName error: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("user Repository FindByName: user not found")
			return nil, customerrors.WrapNotFoundError(err, "user Repository FindByName: user not found")
		}
		return nil, err
	}
	user, err := userdm.ReconstructUser(tempUser.ID, tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, tempUser.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepositoryImpl) FindByUserNames(ctx context.Context, names []string) (map[string]*userdm.User, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)

	if !exists {
		return nil, errors.New("no transaction found in context")
	}
	query := "SELECT * FROM users WHERE name IN (?)"
	var tempUsers []userRequest
	query, args, err := sqlx.In(query, names)
	if err != nil {
		return nil, err
	}

	err = conn.Select(&tempUsers, query, args...)
	if err != nil {
		return nil, customerrors.WrapInternalServerError(err, "FindByNames Select User Repository database error")
	}

	userMap := make(map[string]*userdm.User)
	for _, tempUser := range tempUsers {
		user, err := userdm.ReconstructUser(tempUser.ID, tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, tempUser.CreatedAt)
		if err != nil {
			return nil, customerrors.WrapInternalServerError(err, "FindByNames error converting userRequest to User")
		}

		userMap[tempUser.Name] = user
	}

	return userMap, nil
}

func (repo *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*userdm.User, error) {

	conn, exists := ctx.Value("Conn").(dbOperator)

	if !exists {
		return nil, errors.New("no transaction found in context")
	}
	query := "SELECT * FROM users WHERE email = ?"

	var tempUser userRequest
	err := conn.Get(&tempUser, query, email)
	if err != nil {
		log.Printf("user Repository FindByEmail error: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("user Repository FindByEmail: user not found")
			return nil, customerrors.WrapNotFoundError(err, "user Repository FindByEmail: user not found")
		}
		return nil, err
	}
	user, err := userdm.ReconstructUser(tempUser.ID, tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, tempUser.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepositoryImpl) FindSkillsByUserID(ctx context.Context, userID string) ([]userdm.Skill, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)

	if !exists {
		return nil, errors.New("no transaction found in context")
	}
	query := "SELECT * FROM skills WHERE user_id = ?"

	var tempSkills []skillRequest
	err := conn.Select(&tempSkills, query, userID)
	if err != nil {
		log.Printf("user Repository FindSkillsByUserID error: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("user Repository FindSkillsByUserID: skills not found for userID: %s", userID)
			return nil, customerrors.WrapNotFoundError(err, "user Repository FindSkillsByUserID: skills not found")
		}
		return nil, err
	}

	var skills []userdm.Skill
	for _, tempSkill := range tempSkills {
		skill, err := userdm.ReconstructSkill(tempSkill.ID, tempSkill.TagID, tempSkill.UserID, tempSkill.Evaluation, tempSkill.Years, tempSkill.CreatedAt, tempSkill.UpdatedAt)
		if err != nil {
			return nil, customerrors.WrapInternalServerError(err, "FindSkillsByUserID error converting skillRequest to Skill")
		}
		skills = append(skills, *skill)
	}

	return skills, nil
}
