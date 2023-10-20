mysqldb:
	docker run --name mysqldb -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=todo4 -e MYSQL_USER=todo -e MYSQL_PASSWORD=secret -p 3306:3306 -d mysql:8.1

migrateup:
	migrate -path db/migrations -database "mysql://todo:secret@tcp(localhost:3306)/todo4" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://todo:secret@tcp(localhost:3306)/todo4" -verbose down

server:
	go run main.go

reload:
	@reflex -r '\.go' -s -- sh -c "go run main.go serve"

up:
	docker compose up

clear:
	docker stop $$(docker ps -a -q)
	docker rm $$(docker ps -a -q)
	docker volume rm $$(docker volume ls -q)
	docker rmi $$(docker images -q)
	docker system prune

rm:
	docker container rm -f $$(docker ps -aq)

rmi:
	docker rmi -f $$(docker images -a -q)

prune:
	docker system prune

rm-test:
	docker ps -a --filter "ancestor=monsterup/devcode-unit-test-1" -q | xargs docker rm

test:
	docker ps -a --filter "ancestor=monsterup/devcode-unit-test-1" -q | xargs docker rm
	docker run -e API_URL=http://host.docker.internal:3030 monsterup/devcode-unit-test-1

hapuss:
	docker ps -a --filter "ancestor=slackman/skyshi-todo" -q | xargs docker rm
	docker rmi -f $$(docker images 'skyshi-todo' -a -q)
	docker rmi -f $$(docker images -f "dangling=true" -q)

hapus:
	docker rm $(docker ps -a -q --filter="name=skyshi-todolist-api-1")
	docker rmi $(docker images 'skyshi-todolist_api' -a -q)
	docker compose up

# get variable from env file
include config.env
NAME							= $(APP_NAME)
VERSION 					= $(shell git describe --tags --always)
DOCKER_HUB_REPO		= slackman/skyshi-todo
DOCKER_IMAGE_NAME	= $(NAME):$(VERSION)


build:
	@echo "-> Running $@"
	@docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} --file Dockerfile --tag $(DOCKER_IMAGE_NAME) .

tag:
	@echo "-> Running $@"
	@docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):latest
	@echo $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):latest
	@docker tag $(DOCKER_IMAGE_NAME) $(DOCKER_HUB_REPO):latest
	@docker images

upload:
	@docker push slackman/skyshi-todo

run:
	@docker run -e MYSQL_HOST=172.17.0.2 -e MYSQL_USER=todo -e MYSQL_PASSWORD=secret -e MYSQL_DBNAME=todo4 -p 3030:3030 slackman/skyshi-todo


running: build tag run

.PHONY: mysqldb migrateup migratedown server up clear rm rmi prune hapus push upload