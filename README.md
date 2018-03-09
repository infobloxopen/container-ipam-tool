# container-ipam-tool

Configuration tool for any container environments which uses Infoblox Appliance for their IPAM. This tools pre configure Infoblox Appliance before using it from container environment.
Extensible Attributes in Infoblox Appliance help to manage IPAM objects effectively for any cloud platform. So few EA are needed in Infoblox Appliance for container environment this tool helps to create those EAs

## Prerequisite

A NIOS DDI Appliance with cloud automation License.

You can download a virtual version of the product from the Infoblox Download Center (https://www.infoblox.com/infoblox-download-center). Alternatively, if you are an existing Infoblox customer, you can download it from the support site.

## Configuring Cloud Extensible Attributes using create-ea-defs

If the "Cloud Network Automation" license is activated, then the Cloud Extensible Attributes used by the ipam-plugin can be created using the infoblox/container-ipam-tools docker image. There are 2 ways to do it.

To create EA definition using command line arguments:
```
docker run infoblox/container-ipam-tool:0.0.1 create-ea-defs --debug --cloud-type docker --grid-host 10.120.21.150 --wapi-username=admin --wapi-password=infoblox --wapi-version=2.3
```

To create EA definition using configuration file:
```
docker run -v /etc/infoblox:/etc/infoblox infoblox/container-ipam-tool:0.0.1 create-ea-defs --debug --cloud-type docker --conf-file docker-infoblox.conf
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
```
