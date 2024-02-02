devup:
	docker compose -f dev.docker-compose.yml up -d
devdown:
	docker compose -f dev.docker-compose.yml down --rmi local
resetdev:
	docker compose -f dev.docker-compose.yml down --rmi local
	docker compose -f dev.docker-compose.yml up -d
up:
	docker compose -f docker-compose.yml up -d
down:
	docker compose -f docker-compose.yml down --rmi local

reset:
	docker compose -f docker-compose.yml down --rmi local
	docker compose -f docker-compose.yml up -d