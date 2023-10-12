
# get variable from env file
include config.env
NAME							= $(APP_NAME)
VERSION 					= $(shell git describe --tags --always)
DOCKER_HUB_REPO		= slackman/$(APP_NAME)
DOCKER_IMAGE_NAME	= $(NAME):$(VERSION)

server:
	go run main.go

rm:
	docker container rm $$(docker ps -aq) -f

rmi:
	docker rmi -f $$(docker images -a -q)

prune:
	docker system prune

mysqldb:
	@docker run --name mysqldb -e MYSQL_ROOT_PASSWORD=$(MYSQL_PASSWORD) -e MYSQL_DATABASE=$(MYSQL_DBNAME) -e MYSQL_USER=$(MYSQL_USER) -e MYSQL_PASSWORD=$(MYSQL_PASSWORD) -d  mysql:latest

build:
	@echo "-> Running $@"
	@docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} --file Dockerfile --tag $(DOCKER_IMAGE_NAME) .

push:
	@echo "-> Running $@"
	@docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):latest
	@echo $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):latest
	@docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):latest

upload:
	@docker push $(DOCKER_HUB_REPO)

run:
	@docker run -e MYSQL_HOST=172.17.0.2 -e MYSQL_USER=$(MYSQL_USER) -e MYSQL_PASSWORD=$(MYSQL_PASSWORD) -e MYSQL_DBNAME=$(MYSQL_DBNAME) -p 3030:3030 $(DOCKER_HUB_REPO):latest

.PHONY: server rm rmi prune mysqldb build push upload run