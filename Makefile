.PHONY: test

test: 
	go test -race ./test/...
