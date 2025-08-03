migrate-up:
	go run cmd/migrator/main.go --storage-path=storage/db.sqlite3 --migrations-path=migrations

migrate-down:
	go run cmd/migrator/main.go --storage-path=storage/db.sqlite3 --migrations-path=migrations --down=true

docker:
	docker compose up -d --remove-orphans --build