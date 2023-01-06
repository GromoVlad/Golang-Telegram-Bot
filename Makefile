.PHONY: run
run:
	go build -o cmd/app cmd/main.go && ./cmd/app
migration-status:
	$(MAKE) migration-status -C migration
migration-up:
	$(MAKE) migration-up -C migration
migration-down:
	$(MAKE) migration-down -C migration
migration-create:
	$(MAKE) migration-create -C migration
