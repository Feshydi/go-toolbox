COMPOSE=docker compose -f docker-compose.yml
ENV_FILE ?= --env-file .env
APP_CONTAINER ?= service-go

# -------------- Launching & Terminating --------------

.PHONY: build
build:
	@echo ">>> Building and starting containers..."
	$(COMPOSE) $(ENV_FILE) up --build --remove-orphans -d
	@echo ">>> Containers started"

.PHONY: install
install: build

.PHONY: up
up: install

.PHONY: down
down:
	@echo ">>> Stopping containers..."
	$(COMPOSE) $(ENV_FILE) down --remove-orphans
	@echo ">>> Containers stopped"

.PHONY: clean
clean:
	@echo ">>> Cleaning local environment..."
	$(COMPOSE) $(ENV_FILE) down --volumes --rmi local --remove-orphans
	@echo ">>> Local environment cleaned"

# ---------------------- Tests ----------------------

.PHONY: test-container
test-container:
	docker exec $(APP_CONTAINER) go test -v -race -count=100 ./...

.PHONY: test
test:
	go test -v -race -count=100 ./...

.PHONY: test-cover
test-cover:
	go test -short -count=1 -race -coverprofile=coverage.out $$(go list ./... | grep -v '/mocks$$')
	go tool cover -html=coverage.out
	rm "coverage.out"

.PHONY: generate-mocks
generate-mocks:
	go generate ./...

.PHONY: clean-mocks
clean-mocks:
	find . -type d -name "mocks" -exec rm -rf {} +

.PHONY: regen-mocks
regen-mocks: clean-mocks generate-mocks
