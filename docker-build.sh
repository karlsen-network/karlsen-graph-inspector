#!/bin/bash

set -e

# Verify that all the required environment variables are set
declare -A REQUIRED_VARIABLES
REQUIRED_VARIABLES["KARLSEND_VERSION"]="${KARLSEND_VERSION}"
REQUIRED_VARIABLES["API_ADDRESS"]="${API_ADDRESS}"
REQUIRED_VARIABLES["API_PORT"]="${API_PORT}"

REQUIRED_VARIABLE_NOT_SET=false
for REQUIRED_VARIABLE_NAME in "${!REQUIRED_VARIABLES[@]}"; do
  if [ -z "${REQUIRED_VARIABLES[$REQUIRED_VARIABLE_NAME]}" ]; then
    echo "${REQUIRED_VARIABLE_NAME} is not set";
    REQUIRED_VARIABLE_NOT_SET=true
    fi
done

if [ true = "${REQUIRED_VARIABLE_NOT_SET}" ]; then
  echo
  echo "The following environment variables are required:"
  for REQUIRED_VARIABLE_NAME in "${!REQUIRED_VARIABLES[@]}"; do
    echo "${REQUIRED_VARIABLE_NAME}"
  done
  exit 1
fi

# Build processing
docker build -f processing/Dockerfile -t karlsen-graph-inspector-processing:latest --build-arg KARLSEND_VERSION="${KARLSEND_VERSION}" processing

# Build api
docker build -f api/Dockerfile -t karlsen-graph-inspector-api:latest api

# Build web
REACT_APP_API_ADDRESS="${API_ADDRESS}:${API_PORT}"
docker build -f web/Dockerfile --build-arg REACT_APP_API_ADDRESS="${REACT_APP_API_ADDRESS}" -t karlsen-graph-inspector-web:latest web
