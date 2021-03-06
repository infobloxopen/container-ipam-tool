#!/bin/bash

PLUGIN_DIR="/run/docker/plugins"
GRID_HOST="192.168.124.200"
WAPI_PORT="443"
WAPI_USERNAME=""
WAPI_PASSWORD=""
WAPI_VERSION="2.5"
SSL_VERIFY="false"
CLOUD_TYPE=""


bin/create_ea_defs --grid-host=${GRID_HOST} \
   --wapi-port=${WAPI_PORT} \
   --wapi-username=${WAPI_USERNAME} \
   --wapi-password=${WAPI_PASSWORD} \
   --wapi-version=${WAPI_VERSION} \
   --ssl-verify=${SSL_VERIFY} \
   --cloud-type=${CLOUD_TYPE}
