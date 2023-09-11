start-all:
	docker-compose up -d --remove-orphans --build --force-recreate

start-db:
	docker-compose up -d --remove-orphans --build --force-recreate db

start-service:
	docker-compose up -d --remove-orphans --build --force-recreate library_service

build_and_run_service_locally:
	go build -o app cmd/main/main.go && ./app

test:
	go test -race ./...

lint:
	golangci-lint run
