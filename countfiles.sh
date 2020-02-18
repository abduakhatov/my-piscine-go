files=$(find . -type f -print | wc -l)
dirs=$(find . -type d -print | wc -l)
echo $((files+dirs))
