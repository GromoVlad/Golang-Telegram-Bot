.PHONY: run
run:
	go build -o cmd/app cmd/main.go && ./cmd/app
