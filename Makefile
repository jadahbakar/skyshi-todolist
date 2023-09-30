mysqldb:
	docker run --name mysql81 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=todolist -e MYSQL_USER=todo -e MYSQL_PASSWORD=secret -p 3306:3306 -d mysql:8.1

migrateup:
	migrate -path db/migration -database "mysql://todo:secret@tcp(localhost:3306)/todolist" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://todo:secret@tcp(localhost:3306)/todolist" -verbose down

server:
	go run main.go

clear:
	docker rm $(docker ps -a -q)
	docker rmi $(docker images -a -q)
	docker system prune

.PHONY: mysqldb migrateup migratedown server clear