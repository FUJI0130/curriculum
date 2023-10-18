package middleware

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/config"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	//... [snip]
)

func initializeDB(db *sqlx.DB) error {
	// テスト開始前のDBの状態を初期化（例：特定のテーブルのテストデータを削除）
	tables := []string{
		"contracts",
		"contract_requests",
		"proposals",
		"skills",
		"mentor_recruitments_tags",
		"mentor_recruitments",
		"plans_tags",
		"plans",
		"careers",
		"tags",
		"categories",
		"users", // 他のテーブルも追加することができます
	}

	for _, table := range tables {
		_, err := db.Exec("TRUNCATE " + table + ";")
		if err != nil {
			return err
		}
	}
	return nil
}

func TestTransactionHandling(t *testing.T) {
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
				Name:     "TestNameUser2",
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
		// 追加のテストケース...
	}

	for _, tt := range tests {
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

			// 必要に応じて、DBからデータを取得して、期待する結果と比較する検証を行う

			// テスト終了後のデータベースクリーンアップ
			// if err := initializeDB(db); err != nil {
			// 	t.Fatalf("Failed to cleanup the database: %v", err)
			// }
		})
	}
}