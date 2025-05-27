APP_NAME=lit
TEST_APP=tests/.lit

build:
	go build -o $(APP_NAME) main.go

test:
	go test ./tests
	rm -rf $(TEST_APP)
