include .env
export

env-up:
	docker compose up -d library-postgres

env-down:
	docker compose down library-postgres

run-service:
	docker compose up -d
	
goose-create:
	@if [ -z "$(name)" ]; then \
		echo "There is no name param. Example: make migrate-create name=name_of_migrate"; \
		exit 1; \
	fi;	\

	docker compose run --rm \
	-e GOOSE_COMMAND="create" \
	-e GOOSE_COMMAND_ARG="$(name) sql" \
	goose-migrations

goose-status:
	docker compose run --rm \
	-e GOOSE_COMMAND="status" \
	goose-migrations

goose-up:
	docker compose run --rm \
	-e GOOSE_COMMAND="up" \
	goose-migrations

goose-up-by-one:
	docker compose run --rm \
	-e GOOSE_COMMAND="up-by-one" \
	goose-migrations

goose-down:
	docker compose run --rm \
	-e GOOSE_COMMAND="down" \
	goose-migrations

goose-down-to:
	@if [ -z "$(id)" ]; then \
		echo "There is no migrate id. Example: make goose-down-to id=20170614145246"; \
		exit 1; \
	fi
	docker compose run --rm \
		-e GOOSE_COMMAND="down-to"
		-e GOOSE_COMMAND_ARG="$(id)" \
		goose-migrations

goose-up-to:
	@if [ -z "$(id)" ]; then \
		echo "There is no migrate id. Example: make goose-down-to id=20170614145246"; \
		exit 1; \
	fi
	docker compose run --rm \
		-e GOOSE_COMMAND="up-to"
		-e GOOSE_COMMAND_ARG="$(id)" \
		goose-migrations

goose-full-rollback:
	docker compose run --rm goose-migrations down-to 0 