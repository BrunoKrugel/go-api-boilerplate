all:
	go build ./app/main.go
	go run ./app/main.go

build:
	go build ./app/main.go

run:
	go run ./app/main.go

clean: 
	go mod tidy

update:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

list:
	go install github.com/icholy/gomajor@latest | gomajor list

test:
	go install github.com/mfridman/tparse@latest | go mod tidy
	go test -json -cover ./... | tparse -all -pass

mocks:
	go install github.com/vektra/mockery/v2@latest
	mockery --with-expecter --disable-version-string --all --output internal/mocks

align:
	- fieldalignment -fix ./...		

imports:
	goimports -w .

format:
	go fmt ./...	