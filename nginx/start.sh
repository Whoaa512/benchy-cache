#!/bin/bash
docker build -t nginx-file-cache .
docker run --rm -p 6942:6942 -v $(pwd)/data:/data nginx-file-cache
