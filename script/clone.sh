#!/bin/sh

echo "+ clone ${PLUGIN_CONFIG_PATH} to ${PLUGIN_OUT}"

mkdir ${PLUGIN_OUT}
cd ${PLUGIN_OUT}
git init
git config --global user.name "guoxudong"
git config --global user.email guoxudong@keking.cn
git remote add origin ${PLUGIN_CONFIG_PATH}
git pull origin master

#git clone ${PLUGIN_CONFIG_PATH} ${PLUGIN_OUT}

/bin/drone-git
#echo "configPkg: ${PLUGIN_OUT}" > env.yaml
