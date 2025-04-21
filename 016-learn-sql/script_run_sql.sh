#!/bin/bash

# Check for argument
if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <folder path>"
  exit 1
fi


folder="$1"
printfString=""

for file in "$folder"/*; do
  if [[ -f "$file" ]]; then
    file_name=$(basename "$file")
    printfString+=".mode table\n.read $folder$file_name\\n"
  fi
done

output=$(printf "$printfString" | ./sqlite3.exe)

if [[ -z "$output" ]]; then
  echo "No data found."
  exit 1
fi

printf "$output"

echo ""
echo "Done"