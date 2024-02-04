install:
	go install

build:
	go build -o main main.go

clean:
	rm main

format:
	go fmt

test:
	go test -v ./tests/

server:
	go run main.go