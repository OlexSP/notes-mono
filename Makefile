APP_BIN=app/build/app

.PHONY: build

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./app/cmd/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -rf ./app/build || true

.PHONY: up-local-env
up-local-env: down-local-env
	@docker-compose -f docker-compose.local.yml up -d

.PHONY: down-local-env
down-local-env:
	@docker-compose -f docker-compose.local.yml stop

.PHONY: swagger
swagger:
	#swagger generate spec -o docs/swagger.json
	swag init -g ./app/cmd/main.go -o ./app/docs



.PHONY: migrate
migrate:
	$(APP_BIN) migrate -version $(version)

.PHONY: migrate.up
migrate.up:
	$(APP_BIN) migrate -seq up

.PHONY: migrate.down
migrate.down:
	$(APP_BIN) migrate -seq down