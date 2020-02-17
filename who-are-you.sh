subjects=$(curl -s https://01.alem.school/assets/superhero/all.json | jq -c '.[] | select(.id | contains(70))' | jq -r '.name')

echo \"${subjects}\"
