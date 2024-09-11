
build:
	docker build -t UnionEx .

up:
	docker-compose up -d

down:
	docker-compose down

delete:
	docker rmi UnionEx --force

.PHONY: build up down delete