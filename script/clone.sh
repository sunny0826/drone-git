#!/bin/sh

echo "+ clone ${PLUGIN_CONFIG_PATH} to ${PLUGIN_OUT}"

mkdir ${PLUGIN_OUT}
cd ${PLUGIN_OUT}
git init
git remote add origin ${PLUGIN_CONFIG_PATH}
git pull origin master

#git clone ${PLUGIN_CONFIG_PATH} ${PLUGIN_OUT}

/bin/drone-git
#echo "configPkg: ${PLUGIN_OUT}" > env.yaml
