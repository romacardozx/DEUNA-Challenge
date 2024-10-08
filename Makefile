.PHONY: run down test

build:
	docker compose -f deployments/docker/docker-compose.yml up --build

down:
	docker compose -f deployments/docker/docker-compose.yml down -v

test:
	go clean -testcache && go test -cover -v ./...