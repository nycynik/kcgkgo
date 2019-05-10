#!/usr/bin/env bash

echo starting development

export TZ="America/Los_Angeles"
export PROJECTDIR="/dev/kcgkgo"
export PGSQL_ROOT_PASSWORD="password"
export KC_ADMIN_PASSWORD="password"

docker-compose -f ${PROJECTDIR}/docker/docker-compose.yml up
