# generate_hashes.py

import hashlib
import os

hashes = set()

while len(hashes) < 10000:
    random_data = os.urandom(20)  # Generate 20 random bytes
    hash_value = hashlib.md5(random_data).hexdigest()
    hashes.add(hash_value)

with open('hashes.txt', 'w') as f:
    for hash_value in hashes:
        f.write(hash_value + '\n')
