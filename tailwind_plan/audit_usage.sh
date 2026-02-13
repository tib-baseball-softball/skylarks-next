#!/bin/bash
TARGET_DIR=$1
CLASSES_FILE=$2

# Create a regex from classes (this might be too long for grep)
# Let's do it per file to be safer.

find "$TARGET_DIR" -name "*.svelte" | while read -r file; do
    total_count=0
    # Search for all classes in one pass if possible, or just grep for common prefixes
    # Actually, let's just count how many times "class=" contains one of our target classes roughly.
    # To identify hotspots, even a rough count is enough.

    # Get all class attributes
    classes_in_file=$(grep -o 'class="[^"]*"' "$file" | sed 's/class="//;s/"//' | tr ' ' '\n')
    classes_in_file+=$(grep -o "class='[^']*'" "$file" | sed "s/class='//;s/'//" | tr ' ' '\n')

    # Compare with our list
    matched_count=0
    for c in $classes_in_file; do
        if grep -Fxq "$c" "$CLASSES_FILE"; then
            ((matched_count++))
        fi
    done

    if [ $matched_count -gt 0 ]; then
        echo "$matched_count $file"
    fi
done | sort -nr
