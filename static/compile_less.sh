#!/bin/sh

for file in less/*.less
do
	echo "Processing $file"
	filename=$(basename "$file")
	filename=${filename%.*}
	lessc $file > css/$filename.css
done
