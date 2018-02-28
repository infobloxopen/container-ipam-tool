ifeq ($(DOCKERHUB_ID),)
    	TOOLS_IMAGE_NAME=container-ipam-tools
else
    TOOLS_IMAGE_NAME=${DOCKERHUB_ID}/docker-ipam-tools
endif
RELEASE=1.1.0

.PHONY: clean-plugin
clean-plugin:
	rm -rf ./plugin ./bin
	docker plugin disable ${PLUGIN_NAME}:${RELEASE} || true
	docker plugin rm ${PLUGIN_NAME}:${RELEASE} || true
	docker rm -vf tmp || true
	docker rmi ipam-build-image || true
	docker rmi ${PLUGIN_NAME}:rootfs || true

.PHONY: build-binary
build-binary:
	docker build -t ipam-build-image -f Dockerfile.build .
	docker create --name build-container ipam-build-image
	docker cp build-container:/go/src/github.com/infobloxopen/container-ipam-tool/bin .
	docker rm -vf build-container
	docker rmi ipam-build-image

.PHONY: clean-tools-image
clean-tools-image:
	docker rmi ${TOOLS_IMAGE_NAME}:${RELEASE} || true
	docker rmi ${TOOLS_IMAGE_NAME} || true

.PHONY: build-tools-image
build-tools-image: clean-tools-image build-binary
	docker build -t ${TOOLS_IMAGE_NAME} -f Dockerfile.tools .

.PHONY: push-tools-image
push-tools-image: build-tools-image
	docker tag ${TOOLS_IMAGE_NAME} ${TOOLS_IMAGE_NAME}:${RELEASE}
	docker push ${TOOLS_IMAGE_NAME}:${RELEASE}
