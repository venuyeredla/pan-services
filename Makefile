.PHONY: start-services stop-services 

start-services:
	docker compose start -f ./deployment/docker-compose.yaml

stop-services:
	docker compose stop -f ./deployment/docker-compose.yaml

Run:
	go run main.go

Build:
	go build -o gapp Application.go

BuildIMG:
	docker build . -t goapp:latest

DocRun:
	docker run --name=GoApp -d -p 2024:2024 goapp:latest