PROJECT=studying-clean-architecture-with-golang

RELEASE_MESSAGE=$(shell git log `git describe --tags --abbrev=0`..master --oneline --no-decorate | awk '{printf "%s\\n", $$0}' | sed 's/"/\\"/g')

upgrade-all: upgrade install

upgrade:
	go get -u ./...

clean:
	rm -rf vendor/

install:
	go mod vendor && go mod tidy

build/install:
	go install --ldflags='-w -s -extldflags "-static"' -v -a

lint:
	go list ./... | grep -v casio-api/docs/swagger | xargs -L1 staticcheck -f stylish -fail all -tests

