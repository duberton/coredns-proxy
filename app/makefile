.PHONY: build
build:
	go mod tidy
	CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" --ldflags '-extldflags "-static"' -tags=go_json,nomsgpack -o runlocal/main
	docker build -q -f Dockerfile --build-arg BINARY_PATH=runlocal/main -t localhost/app .