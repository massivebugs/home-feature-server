.PHONY: start
start:
	docker compose -f ./compose.local.yaml up -d

.PHONY: stop
stop:
	docker compose -f ./compose.local.yaml down

.PHONY: restart
restart:
	docker compose -f ./compose.local.yaml restart

.PHONY: restart-web
restart-web:
	docker compose -f ./compose.local.yaml restart web

.PHONY: db-migrate
db-migrate:
	docker build -f ci/docker/Dockerfile.local.migrate -t migrate .
	docker run --rm -t --network=home-feature-server_default -v ./db:/app/db migrate:latest ./main

.PHONY: db-rollback
db-rollback:
	docker build -f ci/docker/Dockerfile.local.migrate -t migrate .
	docker run --rm -t --network=home-feature-server_default -v ./db:/app/db migrate:latest ./main --rollback=1

.PHONY: db-reset
db-reset:
	docker build -f ci/docker/Dockerfile.local.migrate -t migrate .
	docker run --rm -t --network=home-feature-server_default -v ./db:/app/db migrate:latest ./main --reset=1

.PHONY: db-seed
db-seed:
	docker build -f ci/docker/Dockerfile.local.seed -t seed .
	docker run --rm -t --network=home-feature-server_default -v ./db:/app/db seed:latest ./main

.PHONY: sqlc-generate
sqlc-generate:
	docker run --rm -v ./:/app -w /app sqlc/sqlc:1.26.0 -f /app/config/sqlc.yaml generate