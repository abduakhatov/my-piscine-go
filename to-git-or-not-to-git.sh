IP=$(curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"abduakhatov\"}}){id}}"}' | python -m json.tool | grep '\"id\"' | cut -d ':' -f 2)
IP = ${IP#* }
echo "$IP"
