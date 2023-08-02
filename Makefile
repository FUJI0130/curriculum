# ここで指定されている単語と同じファイル名のファイルが仮にあったとしても、コマンドが実行される様になる
.PHONY: help build build-local up down logs ps test

# Makefileを引数無しで実行した場合に実行されるデフォルトのターゲットを指定する　この場合はhelpが実行される
.DEFAULT_GOAL := help

# Dockerのイメージタグを指定する イメージタグよくわからなかったので調べたら、コンテナのバージョン？を管理するために使うものらしい　今回は開発なのでdevとした
DOCKER_TAG := dev

# buildターゲットを定義している。　アプリケーションをデプロイするためのDockerイメージをビルドする作業を自動化する（らしい）
build: ## Build docker image to deploy		docker build -t go_webapp_hands_on --target deploy ./
# ここのイメージ名にあたる所は、docker-composeファイルの中でimageで指定しているものと合わせる方がいいらしい
# この場合　fuji0130/gotodo:${DOCKER_TAG} がイメージ名：タグ名   　バックスラッシュは、まだこのコマンドはこの行で終わりではなく、次の行に続いている　という意味らしい
# --target の部分はマルチステージビルドを行う際の記述 deplpoyはDockerfile内で定義されたビルドステージの名前　
		 docker build -t curriculum:${DOCKER_TAG} \
		  		--target go_dev ./	

# build-localターゲットの指定　ローカル開発環境用のDockerイメージをビルドする作業を自動化する
build-local: ## Build docker image to local development
		docker-compose build --no-cache

up: ## Do docker compose up with hot reload
		docker-compose up 
down: ## Do docker compose down
		docker-compose down
logs: ## Tail docker compose logs
		docker-compose logs -f
ps: ## Check container status
		docker-compose ps
test: ## Execute tests
		go test -race -shuffle=on ./...

dry-migrate: ## Try migration
	mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo --dry-run < ./_tools/mysql/schema.sql

migrate:  ## Execute migration
	mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo < ./_tools/mysql/schema.sql

help: ## Show options
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
				awk 'BEGIN {FS = ":.*?## "}; {printf "\033[036m%-20s\033[0m %s\n", $$1, $$2}'

generate: ## Generate codes
	go generate ./...

