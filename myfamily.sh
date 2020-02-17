#!/bin/bash
curl -s https://01.alem.school/assets/superhero/all.json | jq -c '.[] | select(.id == $id)' --argjson id 1 | jq -r '.connections.relatives'
