APP_NAME = autoappopener
BIN_DIR = tmp
CMD_DIR = ./cmd/

dev:
	air

build:
	go build -o $(APP_NAME).exe $(CMD_DIR)

run:
	go run $(CMD_DIR)

clean:
	rm -rf $(BIN_DIR)/*

deps:
	go mod tidy