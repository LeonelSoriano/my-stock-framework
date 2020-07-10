.PHONY : cov ocov build run

#ENVIRONMENT := $(shell basename $(dir $(abspath $(dir $$PWD))))
BASE_PATH := $(abspath $(lastword ))

hello:
	@echo $(CURDIR)

build:
	#go build -o bin/main main.go

run:
	go run ./src/main.go

test:
	go test -v ./...

cov: clean
	go test -v  -cover -coverprofile=coverage.out  ./... &&\
	go tool cover -html=coverage.out -o coverage.html

ocov: cov
	firefox $(CURDIR)/coverage.html

clean:
	rm -f coverage.out
	rm -f coverage.html

doc:
	godoc -http=:6060
