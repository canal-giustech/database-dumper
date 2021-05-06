dumper:
	@go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o main  ./cmd/dumper.go

restore:
	@go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o main ./cmd/restore.go

clear-on-docker:
	@rm -rf docker/images/dumper/main

clear: clear-on-docker
	@mv main docker/images/dumper/main

docker-dumper:
	@docker build -t giustech/database-dumper ./docker/images/dumper/

docker-restore:
	@docker build -t giustech/database-restore ./docker/images/dumper/


generate-docker-dumper: dumper clear docker-dumper clear-on-docker
	

generate-docker-restore: dumper clear docker-restore clear-on-docker
	

all: build-generate-docker-dumper build-generate-docker-restore