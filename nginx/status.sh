#!/bin/sh

# Compute number of items
items=$(ls -1 /data | wc -l)

# Compute total size
size=$(du -s /data | cut -f1)

echo "{"
echo "  \"size_on_disk\": $size,"
echo "  \"number_of_items\": $items"
echo "}"
