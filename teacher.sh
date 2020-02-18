cd ./the-final-cl-test/mystery
inter_id=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -f 3 -d " " | tr -d "#")
export interview_id=$inter_id
echo $interview_id
cat interviews/interview-$interview_id

# get color and model
color_model_plate=$(grep -oP "Describes it as a\s+\K(\w.+\s).+" interviews/interview-699607)
color=$($color_model_plate | cut -d ' ' -f 1)
model=$($color_model_plate | cut -d ' ' -f 2 | tr -d ',')
plate_numb=$($color_model_plate | grep -o '".*"' | sed 's/"//g')


# get all cars with related params and print
grep -A 5 -E "$(plate_numb | cut).*9" vehicles | grep -A 4 -B 1 'Honda' | grep -A 3 -B 2 'Blue
# todo loop
