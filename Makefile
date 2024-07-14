.PHONY: build
build:
	go build -o cmd/api/main.go .

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: db
db:
	@if [ "$(rollback)" = "1" ]; then \
		ROLLBACK_FLAG=1; \
	else \
		ROLLBACK_FLAG=0; \
	fi; \
	go run cmd/migrate/main.go -rollback=$$ROLLBACK_FLAG;


.PHONY: sqlc-query
sqlc-query:
	@if [ -f database/queries/$(table).sql ]; then \
		echo "Error: database/queries/$(table).sql already exists."; \
		exit 1; \
	else \
		./tool/generate_query_for_sqlc.sh $(table) $(order_by) $(columns) > database/queries/$(table).sql 2>/dev/null && echo "Query generation successful" || echo "Query generation failed"; \
	fi;

.PHONY: sqlc-generate
sqlc-generate:
	./tool/sqlc generate; 

.PHONY: query
query:
	@if [ -f database/queries/$(table).sql ]; then \
		echo "Error: database/queries/$(table).sql already exists."; \
		exit 1; \
	else \
		./tool/generate_query_for_sqlc.sh $(table) $(order_by) $(columns) > database/queries/$(table).sql 2>/dev/null && echo "Query generation successful" || echo "Query generation failed"; \
		./tool/sqlc generate; \
	fi;