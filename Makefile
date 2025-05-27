APP_NAME=lit
TEST_DIRECTORY=tests

build:
	go build -o $(APP_NAME) main.go

test:
	go test ./tests
	rm -rf $(TEST_DIRECTORY)/.lit
