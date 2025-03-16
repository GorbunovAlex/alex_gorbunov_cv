run:
	CONFIG_PATH=config/dev.yaml go run cmd/main.go --env=.env

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/main cmd/main.go

build-run:
	CONFIG_PATH=config/dev.yaml ./bin/main

generate:
	docker build -t alex_gorbunov_cv-app-1:latest . && \
	docker run -d --name alex_gorbunov_cv-app-1-container -v `pwd`:/home/alexandergorbunov/__dev__/pets/alex_gorbunov_cv alex_gorbunov_cv-app-1:latest && \
	sleep 5 && \
	docker exec alex_gorbunov_cv-app-1-container templ generate && \
	docker stop alex_gorbunov_cv-app-1-container && \
	docker rm alex_gorbunov_cv-app-1-container && \
	docker image rm alex_gorbunov_cv-app-1:latest

