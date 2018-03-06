ifeq ($(DOCKERHUB_ID),)
    TOOLS_IMAGE_NAME=container-ipam-tools
else
    TOOLS_IMAGE_NAME=${DOCKERHUB_ID}/container-ipam-tools
endif

.PHONY: build-binary
build-binary:
	docker build -t ${TOOLS_IMAGE_NAME}:${RELEASE} .
	
.PHONY: deps
deps:
	dep ensure

.PHONY: clean
clean:
	docker rmi ${TOOLS_IMAGE_NAME}:${RELEASE} || true
	docker rmi -f ${TOOLS_IMAGE_NAME}:${RELEASE} || true

.PHONY: build-image
build-image: clean build-binary
	
.PHONY: push-image
push-image: build-image
         docker push ${TOOLS_IMAGE_NAME}:${RELEASE}
