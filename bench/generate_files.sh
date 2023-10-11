#!/bin/bash

mkdir -p random_files

# Generate 10 random-sized files
# 1k, 50k, 100k, 500k, 1M, 5M, 10M, 50M, 100M, 250M, 500M
for i in 1 50 100 500 1000 5000 10000 50000 100000 250000 500000; do
    size_in_bytes=$(( $i * 1024 ))
    path=random_files/file_$i
    # check if file exists and is right size
    if [[ -f $path && $(stat -c%s $path) -eq $size_in_bytes ]]; then
        continue
    fi

    head -c $size_in_bytes </dev/urandom > $path
done



