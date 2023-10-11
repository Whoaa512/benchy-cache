import os
import random

hashes = []

with open('hashes.txt', 'r') as f:
    hashes = f.read().splitlines()

files_list = []
# Read the files in the random_files directory
for root, dirs, files in os.walk('random_files'):
    for file in files:
        files_list.append(os.path.join(root, file))


# memoized get size
_files_sizes = {}
def get_size(file_path: str):
    if file_path in _files_sizes:
        return _files_sizes[file_path]
    else:
        size = os.path.getsize(file_path)
        _files_sizes[file_path] = size
        return size

with open('vegeta_targets.txt', 'w') as f:
    for hash in hashes[:100]:
        # Choose a random file for the path
        file_path = random.choice(files_list)
        file_size = get_size(file_path)


        file_name = f"{os.path.basename(file_path)}_{hash}.bin"

        f.write(f"PUT http://127.0.0.1:6942/cache/{file_name}\n")
        f.write(f"Content-Type: application/octet-stream\n")
        f.write(f"Content-Length: {file_size}\n")
        f.write(f"@{file_path}\n\n")  # "@" denotes file input for body in Vegeta

        # Targets for the get requests
        f.write(f"GET http://127.0.0.1:6942/cache/{file_name}\n\n")
