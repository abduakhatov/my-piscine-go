IP=$(curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"abduakhatov\"}}){id}}"}' 2>/dev/null | jq -r '.data.user')

echo ${IP} | jq -c '.[]' | jq -r '.id'
