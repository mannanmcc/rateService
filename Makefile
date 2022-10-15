.PHONY: test run

PROJECT?=rateservice
IMAGE_TAG?=github.com/mannanmcc/rateservice

mod:
	go mod vendor

build-docker:
	docker build --no-cache -t ${IMAGE_TAG} -t github.com/mannanmcc/rateservice .


start-bdd-docker:
	IMAGE=${IMAGE_TAG} docker-compose -f docker-compose.yml down --rmi local
	IMAGE=${IMAGE_TAG} docker-compose -f docker-compose.yml up --build -d --force-recreate
