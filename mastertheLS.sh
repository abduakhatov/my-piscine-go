OUTPUT="$(ls -haAtr --ignore='.??*' --ignore='.[^.]' --ignore='#*' --ignore='.*' --indicator-style=slash --format=commas)"
echo "${OUTPUT}"
