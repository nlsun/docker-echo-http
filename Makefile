DOCKER_REPO = nlsun
DOCKER_NAME = echo-http
DOCKER_IMAGE = $(DOCKER_REPO)/$(DOCKER_NAME)
DOCKER_VERSION = 1.0.0
CGO_ENABLED ?= 0
GOOS ?= linux
GOARCH ?= amd64
GOBUILDFLAGS ?= CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH)


build:
	$(GOBUILDFLAGS) go build -o start .
	docker build -t $(DOCKER_IMAGE) .
	docker tag $(DOCKER_IMAGE):latest $(DOCKER_IMAGE):$(DOCKER_VERSION)

push:
	docker push $(DOCKER_IMAGE):latest
	docker push $(DOCKER_IMAGE):$(DOCKER_VERSION)

run:
	docker run -p 4040:4040 $(DOCKER_IMAGE) /start 4040
