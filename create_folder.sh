#!/bin/bash

# Check for two arguments
if [[ $# -ne 2 ]]; then
  echo "Usage: $0 <source> <destination>"
  exit 1
fi

src="$1"
dist="$2"

# Go up one level and copy
cd .. && cp -r "$src" "$dist"

# Go into the destination if the copy was successful
if [[ -d "$dist" ]]; then
  echo "go to "$dist""
  cd "$dist"
fi

# Find all files ending with .go in the current directory and its subdirectories
find . -name "*.go" -type f -print0 | while IFS= read -r -d $'\0' file; do
  # Check if the file exists (just in case)
  if [[ -f "$file" ]]; then
    # Remove the content of the file by redirecting an empty string to it
    echo "" > "$file"
    echo "Removed content of: $file"
  fi
done

echo "Finished removing content from all .go files."