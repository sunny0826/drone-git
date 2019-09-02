#!/bin/sh

echo "+ clone ${PLUGIN_CONFIG_PATH} to ${PLUGIN_OUT}"

git clone ${PLUGIN_CONFIG_PATH} ${PLUGIN_OUT}
#echo "configPkg: ${PLUGIN_OUT}" > env.yaml
