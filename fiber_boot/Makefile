.PHONY: init run

init:
	docker compose -f ./compose.yml -p fiber up -d

run: init
	go run main.go
