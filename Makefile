DOCKER_REPO = nlsun
DOCKER_NAME = echo-http
DOCKER_VERSION = 1.0.0

DOCKER_IMAGE = ${DOCKER_REPO}/${DOCKER_NAME}

build:
	go build -o start .
	docker build -t $(DOCKER_IMAGE) .
	docker tag $(DOCKER_IMAGE):latest $(DOCKER_IMAGE):$(DOCKER_VERSION)

push:
	docker push $(DOCKER_IMAGE):latest
	docker push $(DOCKER_IMAGE):$(DOCKER_VERSION)
