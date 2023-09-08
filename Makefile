container:
	docker run --name psql-container -p 5432:5432 -e POSTGRES_PASSWORD=qwerty -e POSTGRES_USER=postgres -e POSTGRES_DB=test -d postgres

create-db:
	docker exec -it test-postgres-1 bin/bash

migrate-up:
	migrate -path pkg/db/schema -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" up

migrate-down:
    migrate -path pkg\db\schema -database "postgresql://postgres:qwerty@localhost:5432/postgres?sslmode=disable" down

mock:
	mockgen -source=internal/service/service.go -destination internal/service/mock/mock.go

tests:
	