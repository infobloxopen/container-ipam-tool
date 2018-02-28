Building container-ipam-tools images
========================================================

Building Ipam Tools Image
---------------------
To build the ipam-tools image use the following command:
```
make build-tools-image
```

Pushing Tools Image to Docker Hub
-------------------------------------
The Makefile also includes a build target to push the ipam-tools image to your Docker Hub.
To do that, you need to first setup the following environment variable:
```
export DOCKERHUB_ID="your-docker-hub-id"

```
You can then use the following command to push the image to your Docker Hub:

To Push docker-ipam-plugin image
```
make push-tools-image
```
