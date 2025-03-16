run:
	go run cmd/server/main.go

test:
	go test ./... -v

run_client:
	go run ./cmd/client/main.go

build:
	mkdir -p ./bin
	go build -o ./bin/go-mmuc-server ./cmd/server/main.go

build_client:
	mkdir -p ./bin
	go build -o ./bin/go-mmuc-client ./cmd/client/main.go

docker-build:
	docker build -f ./ci/Dockerfile -t gp-mmuc .
