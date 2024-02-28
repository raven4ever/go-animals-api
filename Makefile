build:
	go build -o bin/animalz-api cmd/api/main.go

run:
	docker compose up -d
	go run cmd/api/main.go

stop:
	docker compose down