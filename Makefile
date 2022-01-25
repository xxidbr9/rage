PKG_NAME = github.com/xxidbr9/rage

build:
	go build -o ./bin/rage ./main.go

install:
	go install ${PKG_NAME}

uninstall:
	go clean -i ${PKG_NAME}

release:
	goreleaser release --rm-dist && npm publish