mysqldb:
	docker run --name mysql81 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=todolist -e MYSQL_USER=todo -e MYSQL_PASSWORD=secret -p 3306:3306 -d mysql:8.1

migrateup:
	migrate -path db/migration -database "mysql://todo:secret@tcp(mysql:3306)/todolist" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://todo:secret@tcp(mysql:3306)/todolist" -verbose down

server:
	go run main.go

up:
	docker compose up

clear:
	docker rm $$(docker ps -a -q) -f
	docker rmi $$(docker images -a -q)
	docker system prune

rm:
	docker container rm $$(docker ps -aq) -f

rmi:
	docker rmi $$(docker images -a -q)

prune:
	docker system prune

hapus:
	docker rm $(docker ps -a -q --filter="name=skyshi-todolist-api-1")
	docker rmi $(docker images 'skyshi-todolist_api' -a -q)
	docker compose up

# get variable from env file
include config.env
NAME							= $(APP_NAME)
VERSION 					= $(shell git describe --tags --always)
DOCKER_HUB_REPO		= slackman/skyshi-todolist
DOCKER_IMAGE_NAME	= $(NAME):$(VERSION)


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
		-t $(DOCKER_IMAGE_NAME) .

run:
	@docker run -e MYSQL_HOST=172.25.0.2 -e MYSQL_USER=todo -e MYSQL_PASSWORD=secret -e MYSQL_DBNAME=todolist -p 8090:3030 slackman/skyshi-todolist


.PHONY: mysqldb migrateup migratedown server up clear rm rmi prune hapus push upload