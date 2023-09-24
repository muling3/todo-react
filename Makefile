migrate-drop:
	migrate -path ./backend/db/migration -database "mysql://root:password@tcp(localhost:3306)/todo" -verbose drop

migrate-down:
	migrate -path ./backend/db/migration -database "mysql://root:password@tcp(localhost:3306)/todo" -verbose down

migrate-up: 
	migrate -path ./backend/db/migration -database "mysql://root:password@tcp(localhost:3306)/todo" -verbose up

create-migration:
	migrate create -ext sql -dir ./backend/db/migration -seq init_schema

start-docker:
	docker start database-mysql