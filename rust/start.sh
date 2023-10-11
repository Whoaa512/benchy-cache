#!/bin/bash

docker build -t rust-file-cache .
docker run --rm -p 6942:6942 -v $(pwd)/data:/data rust-file-cache
