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

func (repo *userRepositoryImpl) FindByUserName(ctx context.Context, name string) (*userdm.UserDomain, error) {
	// ユーザーエンティティを取得
	conn := ctx.Value("Conn")
	log.Printf("FindByUserName: conn: %v", conn)

	user, err := repo.findUserByName(ctx, name) // 既存のメソッドをプライベートメソッドに変更
	if err != nil {
		return nil, err
	}

	// ユーザーIDを使用して関連するスキルを取得
	skills, err := repo.findSkillsByUserID(ctx, user.ID().String())
	if err != nil {
		return nil, err
	}

	// ユーザーIDを使用して関連するキャリアを取得
	careers, err := repo.findCareersByUserID(ctx, user.ID().String())
	if err != nil {
		return nil, err
	}

	// UserDomain オブジェクトを作成
	userDomain := userdm.NewUserDomain(user, skills, careers)
	return userDomain, nil
}

// findUserByName は、名前に基づいてユーザーエンティティを取得するプライベートメソッド
func (repo *userRepositoryImpl) findUserByName(ctx context.Context, name string) (*userdm.User, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)
	log.Printf("findUserByName: conn: %v", conn)
	if !exists {
		return nil, errors.New("コンテキスト内にトランザクションが見つかりません")
	}
	query := "SELECT * FROM users WHERE name = ?"

	var tempUser datamodel.Users
	err := conn.Get(&tempUser, query, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customerrors.WrapNotFoundError(err, "ユーザーが見つかりません")
		}
		return nil, err
	}

	return userdm.ReconstructUser(tempUser.ID, tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, tempUser.CreatedAt)
}

func (repo *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*userdm.User, error) {

	conn, exists := ctx.Value("Conn").(dbOperator)

	if !exists {
		return nil, errors.New("no transaction found in context")
	}
	query := "SELECT * FROM users WHERE email = ?"

	var tempUser datamodel.Users
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

func (repo *userRepositoryImpl) FindByUserID(ctx context.Context, userID string) (*userdm.UserDomain, error) {
	log.Printf("before findUsersByUserID")
	user, err := repo.findUsersByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	log.Printf("before findSkillsByUserID")
	skills, err := repo.findSkillsByUserID(ctx, userID) // このメソッドはプライベートメソッドとする
	if err != nil {
		return nil, err
	}

	log.Printf("before findCareersByUserID")
	careers, err := repo.findCareersByUserID(ctx, userID) // このメソッドはプライベートメソッドとする
	if err != nil {
		return nil, err
	}

	userDomain := userdm.NewUserDomain(user, skills, careers)
	return userDomain, nil
}

func (repo *userRepositoryImpl) Update(ctx context.Context, userDomain *userdm.UserDomain) error {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return errors.New("no transaction found in context")
	}

	// ユーザー情報を更新
	queryUser := "UPDATE users SET name = ?, email = ?, password = ?, profile = ?, updated_at = ? WHERE id = ?"
	if _, err := conn.Exec(queryUser, userDomain.User.Name(), userDomain.User.Email(), userDomain.User.Password(), userDomain.User.Profile(), time.Now(), userDomain.User.ID().String()); err != nil {
		return err
	}

	// スキル情報を更新
	for _, skill := range userDomain.Skills {
		querySkill := "UPDATE skills SET tag_id = ?, evaluation = ?, years = ?, updated_at = ? WHERE id = ?"
		if _, err := conn.Exec(querySkill, skill.TagID().String(), skill.Evaluation().Value(), skill.Year().Value(), time.Now(), skill.ID().String()); err != nil {
			return err
		}
	}

	// キャリア情報を更新
	for _, career := range userDomain.Careers {
		queryCareer := "UPDATE careers SET detail = ?, ad_from = ?, ad_to = ?, updated_at = ? WHERE id = ?"
		if _, err := conn.Exec(queryCareer, career.Detail(), career.AdFrom(), career.AdTo(), time.Now(), career.ID().String()); err != nil {
			return err
		}
	}

	return nil
}

func (repo *userRepositoryImpl) findUsersByUserID(ctx context.Context, userID string) (*userdm.User, error) {
	log.Printf("findUserByUserID: userID: %s", userID)
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		log.Printf("no transaction found in context")
		return nil, errors.New("no transaction found in context")
	}
	log.Printf("before query")
	query := "SELECT * FROM users WHERE id = ?"

	var tempUser datamodel.Users
	log.Printf("before get")
	err := conn.Get(&tempUser, query, userID)
	log.Printf("FindUserByUserID: tempUser: %v", tempUser)
	if err == sql.ErrNoRows {
		log.Printf("sql.ErrNoRows")
		return nil, customerrors.WrapNotFoundError(err, "User not found by userID")
	} else if err != nil {
		return nil, err
	}
	return userdm.ReconstructUser(tempUser.ID, tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, tempUser.CreatedAt)
}

func (repo *userRepositoryImpl) findSkillsByUserID(ctx context.Context, userID string) ([]*userdm.Skill, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return nil, errors.New("no transaction found in context")
	}

	query := "SELECT * FROM skills WHERE user_id = ?"
	var tempSkills []datamodel.Skills
	err := conn.Select(&tempSkills, query, userID)
	if err != nil {
		return nil, err // You should handle not found error and other errors appropriately
	}

	var skills []*userdm.Skill
	for _, tempSkill := range tempSkills {
		skill, err := userdm.ReconstructSkill(tempSkill.ID, tempSkill.TagID, tempSkill.UserID, tempSkill.Evaluation, tempSkill.Years, tempSkill.CreatedAt, tempSkill.UpdatedAt)
		if err != nil {
			return nil, err // Handle the error appropriately
		}
		skills = append(skills, skill)
	}

	return skills, nil
}

func (repo *userRepositoryImpl) findCareersByUserID(ctx context.Context, userID string) ([]*userdm.Career, error) {
	conn, exists := ctx.Value("Conn").(dbOperator)
	if !exists {
		return nil, errors.New("no transaction found in context")
	}

	query := "SELECT * FROM careers WHERE user_id = ?"
	var tempCareers []datamodel.Careers
	err := conn.Select(&tempCareers, query, userID)
	if err != nil {
		return nil, err // You should handle not found error and other errors appropriately
	}

	var careers []*userdm.Career
	for _, tempCareer := range tempCareers {
		career, err := userdm.ReconstructCareer(tempCareer.ID, tempCareer.Detail, tempCareer.AdFrom, tempCareer.AdTo, tempCareer.UserID, tempCareer.CreatedAt, tempCareer.UpdatedAt)
		if err != nil {
			return nil, err // Handle the error appropriately
		}
		careers = append(careers, career)
	}

	return careers, nil
}
