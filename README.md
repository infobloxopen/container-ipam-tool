# container-ipam-tool
Configuration tool for any container environments

## Prerequisite

To ctreate the EA(Extensible Attributes) definition, you need access to the Infoblox DDI product. For evaluation purposes, you can download a
virtual version of the product from the Infoblox Download Center (https://www.infoblox.com/infoblox-download-center)
Alternatively, if you are an existing Infoblox customer, you can download it from the support site.

## Configuring Cloud Extensible Attributes using create-ea-defs tool

If the "Cloud Network Automation" license is activated, then the Cloud Extensible Attributes used by the ipam-plugin
can be defined using the create-ea-defs tool in the infoblox/container-ipam-tools docker image.

To run create-ea-defs:
```
docker run infoblox/container-ipam-tools:0.0.1 create-ea-defs --debug --cloud-type docker --grid-host 10.120.21.150 --wapi-username=admin --wapi-password=infoblox --wapi-version=2.3
```

To use the configuration file for create-ea-defs:
```
docker run -v /etc/infoblox:/etc/infoblox infoblox/container-ipam-tools:0.0.1 create-ea-defs --debug --cloud-type docker --conf-file docker-infoblox.conf
```

Create a file `docker-infoblox.conf` in **`/etc/infoblox/`** directory and add the configuration options in the file.

A sample configuration file looks like this:
```
[grid-config]
grid-host="10.120.21.150"
wapi-port="443"
wapi-username="infoblox"
wapi-password="infoblox"
wapi-version="2.0"
ssl-verify="false"
http-request-timeout=60
http-pool-connections=10

[ipam-config]
global-view="global_view"
global-network-container="172.18.0.0/16"
global-prefix-length=24
local-view="local_view"
local-network-container="192.168.0.0/20,192.169.0.0/22"
local-prefix-length=25
```
