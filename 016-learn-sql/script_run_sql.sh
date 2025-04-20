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
    printfString+=".read $folder$file_name\\n"
  fi
done


printf "$printfString" | ./sqlite3.exe > output.log

echo "Done"