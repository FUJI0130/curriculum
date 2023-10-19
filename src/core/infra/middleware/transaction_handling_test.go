package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/config"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func initializeDB(db *sqlx.DB) error {
	// テスト開始前のDBの状態を初期化（例：特定のテーブルのテストデータを削除）
	tables := []string{
		"careers",
		"skills",
		"users",
		"tags",
		"proposals",
		"plans",
		"categories",
		"plans_tags",
		"mentor_recruitments",
		"mentor_recruitments_tags",
		"contract_requests",
		"contracts",
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
			wantError: true,
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
			wantError: true,
		},
		{
			name: "DBに存在していないスキルが重複",
			input: &userapp.CreateUserRequest{
				Name:     "NonExistentDuplicateSkillUser",
				Email:    "nonExistentDuplicate@gmail.com",
				Password: "password12345",
				Skills: []userapp.SkillRequest{
					{
						TagName:    "NewSkill1",
						Evaluation: 2,
						Years:      26,
					},
					{
						TagName:    "NewSkill1", // 同じスキル名
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
			wantError: true,
		},
		{
			name: "DBに存在しているスキルが重複",
			input: &userapp.CreateUserRequest{
				Name:     "ExistentDuplicateSkillUser",
				Email:    "existentDuplicate@gmail.com",
				Password: "password12345",
				Skills: []userapp.SkillRequest{
					{
						TagName:    "RDBMS",
						Evaluation: 2,
						Years:      26,
					},
					{
						TagName:    "RDBMS", // 既存のスキル名
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
			wantError: true,
		},
		{
			name: "careerに不正な値が入った時にロールバックされるか確認",
			input: &userapp.CreateUserRequest{
				Name:     "TestNameUser2",
				Email:    "validEmail2@gmail.com",
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
						Detail: strings.Repeat("a", 1001), // 1000文字以上（あるいはデータベースの制限を超える文字数）でエラーを起こす仮定
						AdFrom: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						AdTo:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantError: true, // エラーを期待
		},
	}

	// これはテスト関数の外部、最初に一度だけ実行します。
	r := gin.Default()
	r.Use(TransactionHandler(db))

	r.POST("/test-endpoint", func(c *gin.Context) {
		tx, _ := c.Get("tx")
		ctx := context.WithValue(context.Background(), "tx", tx.(*sqlx.Tx))

		var input userapp.CreateUserRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		userRepo := rdbimpl.NewUserRepository(tx.(*sqlx.Tx))
		tagRepo := rdbimpl.NewTagRepository(tx.(*sqlx.Tx))
		existService := userdm.NewExistByNameDomainService(userRepo)
		service := userapp.NewCreateUserAppService(userRepo, tagRepo, existService)

		err := service.ExecWithTransaction(ctx, tx.(*sqlx.Tx), &input)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, "success")
	})

	for _, tt := range tests {
		if selectedTestCase != "" && tt.name != selectedTestCase {
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "正常なリクエスト" {
				if err := initializeDB(db); err != nil {
					t.Fatalf("Failed to initialize the database: %v", err)
				}
			}
			jsonData, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Failed to marshal input: %v", err)
			}
			body := bytes.NewBuffer(jsonData)
			// ダミーエンドポイントを呼び出します。
			resp := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/test-endpoint", body)
			r.ServeHTTP(resp, req)

			if (resp.Code != 200) != tt.wantError {
				t.Errorf("Request failed with code: %d, message: %s", resp.Code, resp.Body.String())
			}
		})
	}

}
