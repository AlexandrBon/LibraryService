start-all:
	docker-compose up -d --remove-orphans --build --force-recreate

start-db:
	docker-compose up -d --remove-orphans --build --force-recreate db

build:
	go build -o app cmd/main/main.go

run:
	./app

test:
	go test -race ./...

lint:
	golangci-lint run
