.PHONY: all

all: # main.go, Dockerfile, go.mod, go.sum
	docker build -t go-backend .
	docker build -t vue-frontend ./frontend

up:
	docker-compose up -d

down:
	docker-compose down

clean: 
	docker image prune


# cd ./frontend && npm run serve
# go run main.go
# docker run -it --rm -p 8888:8888 --name backend go-backend
# docker run -it --rm -p 8080:8080 --name frontend vue-frontend