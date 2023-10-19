package middleware

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/config"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	//... [snip]
)

func initializeDB(db *sqlx.DB) error {
	// テスト開始前のDBの状態を初期化（例：特定のテーブルのテストデータを削除）
	tables := []string{
		"users",
		"tags",
		"careers",
		"proposals",
		"plans",
		"categories",
		"plans_tags",
		"mentor_recruitments",
		"mentor_recruitments_tags",
		"contract_requests",
		"contracts",
		"skills",
	}

	for _, table := range tables {
		// _, err := db.Exec("TRUNCATE " + table + ";")
		_, err := db.Exec("DELETE FROM " + table + ";")

		if err != nil {
			return err
		}
	}
	return nil
}

func TestTransactionHandling(t *testing.T) {
	fmt.Println("config.Env is : ", config.Env)
	selectedTestCase := os.Getenv("SELECTED_TEST_CASE")
	// DBの接続情報を設定ファイルから取得
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		config.Env.TestDbUser,
		config.Env.TestDbPassword,
		config.Env.TestDbHost,
		config.Env.TestDbPort,
		config.Env.TestDbName))

	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	tests := []struct {
		name      string
		input     *userapp.CreateUserRequest
		wantError bool
	}{
		{
			name: "正常なリクエスト",
			input: &userapp.CreateUserRequest{
				Name:     "TestNameUser",
				Email:    "validEmail@gmail.com",
				Password: "password12345",
				Skills: []userapp.SkillRequest{
					{
						TagName:    "RDBMS",
						Evaluation: 2,
						Years:      26,
					},
					{
						TagName:    "AWS",
						Evaluation: 4,
						Years:      26,
					},
					{
						TagName:    "docker",
						Evaluation: 5,
						Years:      26,
					},
				},
				Profile: "Software Developer",
				Careers: []userapp.CareersRequest{
					{
						Detail: "Worked at XYZ Corp",
						AdFrom: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						AdTo:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantError: false,
		},
		{
			name: "無効なEメールアドレス",
			input: &userapp.CreateUserRequest{
				Name:     "InvalidEmailUser",
				Email:    "invalidEmail",
				Password: "password12345",
				Skills: []userapp.SkillRequest{
					{
						TagName:    "RDBMS",
						Evaluation: 2,
						Years:      26,
					},
					{
						TagName:    "AWS",
						Evaluation: 4,
						Years:      26,
					},
					{
						TagName:    "docker",
						Evaluation: 5,
						Years:      26,
					},
				},
				Profile: "Software Developer",
				Careers: []userapp.CareersRequest{
					{
						Detail: "Worked at XYZ Corp",
						AdFrom: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						AdTo:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantError: false,
		},
		{
			name: "既に存在するEメールアドレス",
			input: &userapp.CreateUserRequest{
				Name:     "ExistingEmailUser",
				Email:    "validEmail@gmail.com", // 既存のメールアドレスと同じ
				Password: "password12345",
				Skills: []userapp.SkillRequest{
					{
						TagName:    "RDBMS",
						Evaluation: 2,
						Years:      26,
					},
					{
						TagName:    "AWS",
						Evaluation: 4,
						Years:      26,
					},
					{
						TagName:    "docker",
						Evaluation: 5,
						Years:      26,
					},
				},
				Profile: "Software Developer",
				Careers: []userapp.CareersRequest{
					{
						Detail: "Worked at XYZ Corp",
						AdFrom: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						AdTo:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantError: false,
		},
		{
			name: "パスワードが短すぎる",
			input: &userapp.CreateUserRequest{
				Name:     "ShortPasswordUser",
				Email:    "shortPasswordEmail@gmail.com",
				Password: "pass",
				Skills: []userapp.SkillRequest{
					{
						TagName:    "RDBMS",
						Evaluation: 2,
						Years:      26,
					},
					{
						TagName:    "AWS",
						Evaluation: 4,
						Years:      26,
					},
					{
						TagName:    "docker",
						Evaluation: 5,
						Years:      26,
					},
				},
				Profile: "Software Developer",
				Careers: []userapp.CareersRequest{
					{
						Detail: "Worked at XYZ Corp",
						AdFrom: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						AdTo:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantError: false,
		},
		{
			name: "スキルが重複",
			input: &userapp.CreateUserRequest{
				Name:     "DuplicatedSkillUser",
				Email:    "duplicatedSkill@gmail.com",
				Password: "password12345",
				Skills: []userapp.SkillRequest{
					{
						TagName:    "RDBMS",
						Evaluation: 2,
						Years:      26,
					},
					{
						TagName:    "RDBMS", // 同じスキル名
						Evaluation: 4,
						Years:      26,
					},
				},
				Profile: "Software Developer",
				Careers: []userapp.CareersRequest{
					{
						Detail: "Worked at XYZ Corp",
						AdFrom: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						AdTo:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {

		if selectedTestCase != "" && tt.name != selectedTestCase { // 環境変数が設定されていて、名前が一致しない場合はスキップ
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			if err := initializeDB(db); err != nil {
				t.Fatalf("Failed to initialize the database: %v", err)
			}

			tx, err := db.Beginx()
			if err != nil {
				t.Fatalf("Failed to begin transaction: %v", err)
			}

			c := &gin.Context{} // テスト用のGinコンテキストを作成
			c.Set("tx", tx)     // コンテキストにトランザクションを設定

			defer func() {
				if p := recover(); p != nil {
					tx.Rollback()
					t.Fatalf("Test panicked: %v", p)
				} else if t.Failed() {
					tx.Rollback()
				} else {
					tx.Commit()
				}
			}()

			// リポジトリとドメインサービスのセットアップ
			userRepo := rdbimpl.NewUserRepository(tx) // トランザクションを考慮してtxを使用
			tagRepo := rdbimpl.NewTagRepository(tx)   // 同上
			existService := userdm.NewExistByNameDomainService(userRepo)
			service := userapp.NewCreateUserAppService(userRepo, tagRepo, existService)

			err = service.Exec(context.Background(), tt.input) // トランザクションを含むExecメソッドの呼び出し
			if (err != nil) != tt.wantError {
				t.Errorf("Error = %v, wantError %v", err, tt.wantError)
			}

		})
	}
}
