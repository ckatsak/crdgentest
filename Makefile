DOCKER_BUILD_FLAGS := --force-rm --no-cache
DOCKER_REPO := ckatsak
DISTTATE_BIN := disttate_static

.PHONY: all disttate


all: disttate


disttate:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a \
		    --ldflags '-s -w -extldflags "-static"' -tags netgo \
		    --installsuffix netgo -o deployment/$(DISTTATE_BIN) \
		    ./cmd/disttate
	-docker rmi $(DOCKER_REPO)/disttate
	docker build $(DOCKER_BUILD_FLAGS) -t $(DOCKER_REPO)/disttate \
		-f deployment/Dockerfile $(CURDIR)/deployment
	rm -v deployment/$(DISTTATE_BIN)
	docker push $(DOCKER_REPO)/disttate
