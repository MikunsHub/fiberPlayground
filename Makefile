build:
	go build -o server main.go

run: build
	./server

watch:
	$(shell go env GOPATH)/bin/reflex -s -r '\.go$$' make run
