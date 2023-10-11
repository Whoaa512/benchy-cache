#!/bin/bash -e


bash bench/generate_files.sh

python3 becnh/generate_hashes.py
python3 becnh/generate_vegeta_targets.py

echo "Running benchmarks..."

# https://github.com/tsenart/vegeta
# TODO: make the endpoints configurable
cat vegeta_targets.txt | vegeta attack -duration=30s -rate=100 | tee results.bin | vegeta report
