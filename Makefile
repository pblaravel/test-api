setup:
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./cmd/server/main.go -o ./docs
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files

docs-build:
	swag init -g ./cmd/server/main.go -o ./docs


update-server:
	git pull
	docker compose up -d --build
	docker logs -f dockerBack

build:
	docker compose build --no-cache

certgen:
	echo "creating server.key"
	openssl genrsa -out server.key 2048
	openssl ecparam -genkey -name secp384r1 -out server.key
	echo "creating server.crt"
	openssl req -new -x509 -sha256 -key server.key -out server.crt -batch -days 3650

run-local:
	docker start dockerPostgres
	docker start dockerRedis
	docker start dockerMongo
	docker start dockerRabbit
	export REDIS_HOST=localhost
	export POSTGRES_DB=go_app_dev
	export POSTGRES_USER=docker
	export POSTGRES_PASSWORD=password
	export POSTGRES_PORT=5435
	export JWT_SECRET_KEY=ObL89O3nOSSEj6tbdHako0cXtPErzBUfq8l8o
	export API_SECRET_KEY=cJGZ8L1sDcPezjOy1zacPJZxzZxrPObm2Ggs1U0V
	export POSTGRES_HOST=localhost
	go run cmd/server/main.go

up:
	docker compose up -d --build

down:
	docker compose down

restart:
	docker compose restart

clean:
	docker stop go-space-app-backend
	docker stop dockerPostgres
	docker rm go-space-app-backend
	docker rm dockerPostgres
	docker rm dockerRedis
	docker image rm golang-space-app-backend
	rm -rf .dbdata
