cd ./the-final-cl-test/mystery
inter_id=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -f 3 -d " " | tr -d "#")
export interview_id=$inter_id
echo $interview_id
cat interviews/interview-699607

