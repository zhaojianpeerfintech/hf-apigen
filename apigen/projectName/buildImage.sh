#!/bin/bash
echo ${DOCKERIMAGE}
docker build -t ${DOCKERIMAGE} .
