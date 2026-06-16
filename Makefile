.PHONY: dev api-test api-build frontend-build logs down

dev:
	docker compose up --build

api-test:
	cd backend && go test ./...

api-build:
	cd backend && go build ./cmd/server

frontend-build:
	docker build -t bringit-frontend ./frontend

logs:
	docker compose logs -f

down:
	docker compose down
