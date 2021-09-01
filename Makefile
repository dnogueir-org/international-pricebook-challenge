install:
	docker-compose up -d

tests:
	docker exec -it productserver go test ./...

run:
	docker exec -it productserver go run main.go
	
stop:
	docker-compose down