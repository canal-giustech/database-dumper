dumper:
	@go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o main  ./cmd/dumper.go

restore:
	@go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o main ./cmd/restore.go

clear:
	@rm -rf docker/images/dumper/main
	@mv main docker/images/dumper/main

build-generate-docker-dumper: dumper clear
	@docker build -t go-dumper:latest ./docker/images/dumper/

build-generate-docker-restore: dumper clear
	@docker build -t go-restore:latest ./docker/images/dumper/

build-all: build-generate-docker-dumper build-generate-docker-restore