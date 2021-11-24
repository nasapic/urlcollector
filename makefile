DEV_TAG=dev
STG_TAG=stage
PROD_TAG=v0.0.1
IMAGE_NAME=nasapic-urlcollector

# Names
BINARY_NAME=urlcollector
BINARY_UNIX=$(BINARY_NAME)_unix

# Misc
CONTAINER_NAME=nasapi_urlcollector

# Accounts
DOCKER_ACCOUNT=adrianpksw

.PHONY: mod-tidy
mod-tidy:
	GOPRIVATE=gitlab.com/QWRyaWFuIEdvR29BcHBzIE5BU0E/* go mod tidy

.PHONY: get-deps
get-deps:
	make mod-tidy

.PHONY: build
build:
	go build -o ./bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME).go

.PHONY: build-linux
build-linux:
	CGOENABLED=0 GOOS=linux GOARCH=amd64; go build -o ./bin/$(BINARY_UNIX) ./cmd/$(BINARY_NAME).go

.PHONY: format
format:
	gofmt -s -w .

.PHONY: run
run:
	make build
	make setapikey
	./bin/urlcollector -json-api-port 8080 -nasa-api-key-envar API_KEY -logging-level all

.PHONY: test
test:
	make -f makefile.test test-selected

.PHONY: setapikey
setsampleapikey:
	export API_KEY="DEMO_KEY"

.PHONY: docker-build-dev
docker-build-dev:
	make build
	docker login
	docker build -f deployments/docker/Dockerfile.dev -t ${DOCKER_ACCOUNT}/$(IMAGE_NAME):$(DEV_TAG) .
	docker push ${DOCKER_ACCOUNT}/$(IMAGE_NAME):$(DEV_TAG)

.PHONY: docker-build-stg
docker-build-stg:
	make build
	docker login
	@echo "Building Docker image ${DOCKER_ACCOUNT}/$(IMAGE_NAME):$(STG_TAG)"
	docker build -t ${DOCKER_ACCOUNT}/$(IMAGE_NAME):$(STG_TAG) .
	docker push ${DOCKER_ACCOUNT}/$(IMAGE_NAME):$(STG_TAG)

.PHONY: docker-build-prod
docker-build-prod:
	make build
	docker login
	docker build -t ${DOCKER_ACCOUNT}/$(IMAGE_NAME):$(PROD_TAG) .
	docker push ${DOCKER_ACCOUNT}/$(IMAGE_NAME):$(PROD_TAG)

.PHONY: connect-prod
connect-prod:
	gcloud beta container clusters get-credentials ${GC_PROD_CLUSTER} --region ${GC_REGION} --project ${GC_PROD_PROJECT}

.PHONY: install-prod
install-prod:
	make connect-prod
	helm install --name $(IMAGE_NAME) -f ./deployments/helm/values-prod.yaml ./deployments/helm

.PHONY: docker-run-dev
docker-run-dev:
	docker run -d --name ${CONTAINER_NAME} -p 8080:8080 -it ${DOCKER_ACCOUNT}/$(IMAGE_NAME):$(DEV_TAG)

.PHONY: docker-start-dev
docker-start-dev:
	docker start ${CONTAINER_NAME}

.PHONY: docker-stop-dev
docker-stop-dev:
	docker stop ${CONTAINER_NAME}

spacer:
	@echo "\n"

port8080-list:
	lsof -i :8080
