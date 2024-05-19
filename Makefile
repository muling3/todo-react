migrate-drop:
	migrate -path ./backend/db/migration -database "mysql://root:password@tcp(localhost:3306)/todo" -verbose drop

migrate-down:
	migrate -path ./backend/db/migration -database "mysql://root:password@tcp(localhost:3306)/todo" -verbose down

migrate-up: 
	migrate -path ./backend/db/migration -database "mysql://root:password@tcp(localhost:3306)/todo" -verbose up

create-migration:
	migrate create -ext sql -dir ./backend/db/migration -seq init_schema

generate: 
	cd backend && sqlc generate

start-db:
	docker start database-mysql

backend:
	cd backend && go run main.go

frontend:
	npm run start

.PHONY: migrate-drop migrate-down migrate-up create-migration generate start-db frontend backend