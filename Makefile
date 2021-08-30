.PHONY: test up build run stop down docker-login image-push

test:
	docker-compose -f docker-compose.yml up --build -d -V --force-recreate --remove-orphans
	go test -race -count=1 -p 1 -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out ./...
	docker-compose -f docker-compose.yml rm -sfv

up:
	docker-compose -f docker-compose.yml up -d

build:
	docker-compose -f docker-compose.yml up -d --force-recreate --remove-orphans --build

stop:
	docker-compose stop

down:
	docker-compose -f docker-compose.yml kill
	docker-compose -f docker-compose.yml rm -sfv

docker-login:
	docker login -u=${DOCKER_USERNAME} -p=${DOCKER_PASSWORD} ${DOCKER_REGISTRY}

image-push: docker-login
	docker push ${IMAGE_NAME}:${VERSION}
	docker push ${IMAGE_NAME}:latest