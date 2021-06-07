#!/bin/bash
echo ${DOCKERIMAGE}
docker build -t peerfintech/hf-apigen .
