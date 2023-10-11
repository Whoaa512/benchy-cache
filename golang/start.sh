#!/bin/bash

docker build -t golang-file-cache .
docker run --rm -p 6942:6942 -v $(pwd)/data:/data golang-file-cache
