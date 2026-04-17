APP_NAME = autoappopener
BIN_DIR = tmp
CMD_DIR = ./cmd/

# tells Make these are not files
.PHONY: dev build run clean deps

dev:
	air

build:
	go build -o build/$(APP_NAME).exe $(CMD_DIR)

run:
	go run $(CMD_DIR)

clean:
	rm -rf $(BIN_DIR)/*

deps:
	go mod tidy