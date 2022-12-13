OS=linux
ARCH=amd64

run:
	docker-compose up

db-run:
	docker-compose --profile db up

app-run:
	docker-compose --profile app up

app-run-build:
	docker-compose --profile app up --build

stop:
	docker-compose down