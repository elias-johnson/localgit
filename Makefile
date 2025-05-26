APP_NAME=localgit

build:
	go build -o $(APP_NAME) main.go

clean:
	rm -f $(APP_NAME)

fmt:
	go fmt ./...
