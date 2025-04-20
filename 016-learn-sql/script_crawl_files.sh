#!/bin/bash

# Check if exactly 2 arguments are provided
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <URL> <output_folder>"
  exit 1
fi

# Parameters
URL="$1"
OUTPUT_FOLDER="$2"

# Array of target values
TARGETS=("001_up.sql" "002_main.sql" "003_test.sql") # Modify this array with your targets

# Create output folder if it doesn't exist
mkdir -p "$OUTPUT_FOLDER" || {
  echo "Failed to create or access folder: $OUTPUT_FOLDER"
  exit 1
}

# Fetch HTML and extract the script tag with id="__NUXT_DATA__" into a variable
NUXT_DATA=$(curl -s "$URL" | \
  grep -oP '<script[^>]*id="__NUXT_DATA__"[^>]*>.*?</script>' | \
  sed -E 's/.*<script[^>]*id="__NUXT_DATA__"[^>]*>(.*?)<\/script>/\1/')

# Check if NUXT_DATA is empty
if [ -z "$NUXT_DATA" ]; then
  echo "Failed to extract __NUXT_DATA__ content from $URL"
  exit 1
fi

# Process each target
for TARGET in "${TARGETS[@]}"; do
  # Find the index of the target value in the JSON array
  INDEX=$(echo "$NUXT_DATA" | jq -r --arg target "$TARGET" 'to_entries | map(select(.value == $target)) | .[0].key // "not found"')

  if [ "$INDEX" == "not found" ]; then
    echo "Value '$TARGET' not found in the JSON array"
    continue
  fi

  # Calculate the index of the next item (index + 1)
  NEXT_INDEX=$((INDEX + 1))

  # Extract the content at the next index
  NEXT_ITEM=$(echo "$NUXT_DATA" | jq -r --argjson idx "$NEXT_INDEX" '.[$idx] // "not found"')

  if [ "$NEXT_ITEM" == "not found" ]; then
    echo "No item found at index $NEXT_INDEX for target '$TARGET'"
    continue
  fi

  # Create a file named after the target in the output folder and write the next item's content
  OUTPUT_FILE="$OUTPUT_FOLDER/$TARGET"
  echo "$NEXT_ITEM" > "$OUTPUT_FILE"

  # Check if the file was created successfully
  if [ -f "$OUTPUT_FILE" ]; then
    echo "File '$OUTPUT_FILE' created with content from index $NEXT_INDEX"
  else
    echo "Failed to create file '$OUTPUT_FILE'"
  fi
done