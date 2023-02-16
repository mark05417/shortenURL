.PHONY: all

all: # main.go, Dockerfile, go.mod, go.sum
	docker build -t go-server .

run:
	docker build -t go-server .
	docker run -it --rm -p 8888:8888 --name app go-server

clean: 
	docker image prune