.PHONY: lint test clean cov-report dist

export GO111MODULE=on

default: lint test

lint: 
	staticcheck ./...

test: dist
	go test -v -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

cov-report:
	go tool cover -html=coverage.out -o coverage.html

dist:
	mkdir -p dist/plugins-local/src/github.com/softwaremastermind/defaultcspheader
	cp defaultcsp.go dist/plugins-local/src/github.com/softwaremastermind/defaultcspheader/defaultcsp.go
	cp .traefik.yml dist/plugins-local/src/github.com/softwaremastermind/defaultcspheader/.traefik.yml
	tar -czf defaultcsp-plugin.tar.gz -C dist .

clean:
	rm -f coverage.html
	rm -f coverage.out
	rm -rf  dist
	rm defaultcsp-plugin.tar.gz