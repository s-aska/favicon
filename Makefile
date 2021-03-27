.DEFAULT_GOAL := help

PROJECT_NAME="nl-favicon"

create: ## GCPプロジェクト作成
	gcloud app create --region=asia-northeast1 --project $(PROJECT_NAME)

installdeps: ## go mod vendor
	go mod vendor

dev-app: ## go run app/main.go
	GAE_SERVICE=default go run app/main.go

deploy-app: ## gcloud app deploy
	gcloud app deploy --project $(PROJECT_NAME) app.yaml

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help
