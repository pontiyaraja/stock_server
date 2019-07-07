SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD})
install:
	cd cmd
	go get
	go install
run:
	go run cmd/main.go
fmt:
	gofmt -s -l $(TARGET)
test:
	go test stockadapter/worldtradingconnector_test.go
	echo "done"