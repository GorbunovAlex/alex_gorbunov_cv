run:
	CONFIG_PATH=config/dev.yaml go run cmd/main.go --env=.env

generate:
	docker build -t alex_gorbunov_cv-app-1:latest . && \
	docker run -d --name alex_gorbunov_cv-app-1-container -v `pwd`:/home/alexandergorbunov/__dev__/pets/alex_gorbunov_cv alex_gorbunov_cv-app-1:latest && \
	sleep 5 && \
	docker exec alex_gorbunov_cv-app-1-container templ generate && \
	docker stop alex_gorbunov_cv-app-1-container && \
	docker rm alex_gorbunov_cv-app-1-container && \
	docker image rm alex_gorbunov_cv-app-1:latest

