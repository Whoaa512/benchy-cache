# benchy-cache

Spec:
  - Simple file cache webserver exposed on port 6942
  - GET /cache/$path
    - Returns previously saved file or 404
  - PUT /cache/$path
    - Saves to configured disk location, defaults to ./data
      - e.g. ./data/$path
  - GET /status
    - Returns json of current status as defined below
    - Status
      - current size on disk, based on the size of files stored
      - number of items stored




