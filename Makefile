.PHONY: test run

PROJECT?=rateservice
IMAGE_TAG?=github.com/mannanmcc/rateservice

mod:
	go mod vendor

build-docker:
	docker build --no-cache -t ${IMAGE_TAG} -t github.com/mannanmcc/rateservice .

stop-bdd-docker:
	IMAGE=${IMAGE_TAG} docker-compose -f ./test/docker-compose.yml down

start-bdd-docker: build-docker
	IMAGE=${IMAGE_TAG} docker-compose -f ./test/docker-compose.yml down --rmi local
	IMAGE=${IMAGE_TAG} docker-compose -f ./test/docker-compose.yml up --build -d --force-recreate

run-bdd-tests:
	IMAGE=${IMAGE_TAG} docker-compose -f ./test/docker-compose.yml exec -T tests ginkgo -mod vendor -r --randomize-all --trace ./test/...