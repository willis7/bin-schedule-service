.PHONY: build run

NAME = trash-collection-svc

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

install:
	glide install

docker-build:
	docker build -no-cache -rm -t $(NAME) .

docker-run:
	docker run -it --rm --name $(NAME) -p 80:80 $(NAME)

docker-rm:
	docker rm $(NAME)

docker-rmi:
	docker rmi $(NAME)

default: docker-build
