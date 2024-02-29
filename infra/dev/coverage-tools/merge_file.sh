#!/bin/bash

# Find all coverage files in the current directory
coverage_files=$(find . -name "*_coverage.out")

# Ensure at least one coverage file is found
if [ -z "$coverage_files" ]; then
    echo "Error: No coverage files found."
    exit 1
fi

# Merge the coverage files into a temporary file
temp_file="temp_coverage.out"
touch "$temp_file"

first_file=true

for file in $coverage_files; do
    if [ "$first_file" = true ]; then
        cat "$file" >> "$temp_file"
        first_file=false
    else
        tail -n +2 "$file" >> "$temp_file"
    fi
done

# Clean up
mv "$temp_file" coverage.out

echo "Coverage files merged successfully with 'mode:' lines handled."
