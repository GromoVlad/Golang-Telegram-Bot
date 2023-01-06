.PHONY: run
migration-status:
	goose postgres 'postgres://postgres:@localhost:5432/gin_tonic' status
migration-up:
	goose postgres 'postgres://postgres:@localhost:5432/gin_tonic' up
migration-down:
	goose postgres 'postgres://postgres:@localhost:5432/gin_tonic' down
migration-create:
	goose postgres 'postgres://postgres:@localhost:5432/gin_tonic' create $(name) sql