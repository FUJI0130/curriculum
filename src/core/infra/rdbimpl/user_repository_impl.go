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

type userRepositoryImpl struct {
	Conn Queryer
	// db   *sqlx.DB//使ってないのでは
}

type userRequest struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Profile   string    `db:"profile"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUserRepository(conn Queryer) userdm.UserRepository {
	return &userRepositoryImpl{Conn: conn}
}
func (repo *userRepositoryImpl) Store(ctx context.Context, userdomain *userdm.UserDomain) error {

	queryUser := "INSERT INTO users (id, name, email, password, profile, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := repo.Conn.Exec(queryUser, userdomain.User.ID().String(), userdomain.User.Name(), userdomain.User.Email(), userdomain.User.Password(), userdomain.User.Profile(), userdomain.User.CreatedAt().DateTime(), userdomain.User.UpdatedAt().DateTime())
	if err != nil {
		return err
	}

	for _, skill := range userdomain.Skills {
		querySkill := "INSERT INTO skills (id,tag_id,user_id,created_at,updated_at, evaluation, years) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err = repo.Conn.Exec(querySkill, skill.ID().String(), skill.TagID().String(), userdomain.User.ID().String(), skill.CreatedAt().DateTime(), skill.UpdatedAt().DateTime(), skill.Evaluation().Value(), skill.Year().Value())
		if err != nil {
			return err
		}
	}

	for _, career := range userdomain.Careers {
		queryCareer := "INSERT INTO careers (id,user_id, detail, ad_from, ad_to, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err = repo.Conn.Exec(queryCareer, career.ID().String(), career.UserID().String(), career.Detail(), career.AdFrom(), career.AdTo(), career.CreatedAt().DateTime(), career.UpdatedAt().DateTime())
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *userRepositoryImpl) StoreWithTransaction(ctx context.Context, userdomain *userdm.UserDomain) error {

	conn, exists := ctx.Value("Conn").(Transaction)

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
			return customerrors.WrapInternalServerError(err, "Failed to store career")
		}
	}
	return nil
}

func (repo *userRepositoryImpl) FindByName(ctx context.Context, name string) (*userdm.User, error) {
	query := "SELECT * FROM users WHERE name = ?"

	var tempUser userRequest
	err := repo.Conn.Get(&tempUser, query, name)
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

func (repo *userRepositoryImpl) FindByNames(ctx context.Context, names []string) (map[string]*userdm.User, error) {
	query := "SELECT * FROM users WHERE name IN (?)"
	var tempUsers []userRequest
	query, args, err := sqlx.In(query, names)
	if err != nil {
		return nil, err
	}

	err = repo.Conn.Select(&tempUsers, query, args...)
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
