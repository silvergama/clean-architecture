PROJECT=studying-clean-architecture-with-golang

RELEASE_MESSAGE=$(shell git log `git describe --tags --abbrev=0`..master --oneline --no-decorate | awk '{printf "%s\\n", $$0}' | sed 's/"/\\"/g')

upgrade-all: upgrade install

upgrade:
	go get -u ./...

clean:
	rm -rf vendor/

install-tools:
	cat tools/tools.go | grep "_" | awk -F '"' '{print $$2}' | xargs -L1 go get -u

install: install-tools
	go mod vendor && go mod tidy

build/install:
	go install --ldflags='-w -s -extldflags "-static"' -v -a

lint:
	go list ./... | xargs -L1 staticcheck -f stylish -fail all -tests

test:
	go test -v ./...
