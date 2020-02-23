cd mystery

# GENDER and HEIGHT 
scene=$(grep "CLUE" crimescene | head -n 1)
gender=$( echo "$scene" | cut -d ' ' -f18 | tr -d ,)
height=$( echo "$scene" | cut -d ' ' -f21 | tr -d \'.)
# echo "GENDER: $gender"
# echo "HEIGHT: $height"

# PEOPLE INFO
grep "Annabel" people | while read -r people ; do
    w_full_name=$( echo "$people" | awk '{ print $1, $2 }')
    w_gender=$( echo "$people" | awk '{ print $3 }')
    w_age=$( echo "$people" | awk '{ print $4 }')
    w_adress=$( echo "$people" | awk  '{ printf("%s_%s",$5,$6) }' | tr -d ',')
    w_line=$( echo "$people" | awk '{ print $8 }')
    # echo "$w_full_name, $w_gender, $w_age, $w_adress, $w_line"

    # find INTERVIEW by ADDRESS
    inter_id=$(head -n $w_line streets/$w_adress | tail -n 1 | cut -f 2 -d '#' )
    # echo $inter_id

    interview=$(cat interviews/interview-$inter_id | grep -oP '(?<= Describes it as )\w.*' | tr -d '\",')
    if [ -z "$interview" ]
    then
        # echo "$inter_id continue"
        continue        
    fi
    car_color=$( echo "$interview" | cut -d ' ' -f2)
    car_model=$( echo "$interview" | cut -d ' ' -f3)
    car_plate_start=$( echo "$interview" | cut -d ' ' -f11)
    car_plate_end=$( echo "$interview" | cut -d ' ' -f15)
    # echo "$car_color, $car_model, $car_plate_start, $car_plate_end"

    # search for VEHICLE
    # vehicles=$(grep -A 5 -E "License Plate $(car_plate_start | cut).*9$" vehicles | grep -A 4 -B 1 'Honda' | grep -A 3 -B 2 'Blue' | grep -A 1 -B 4 'Height: 6')






done

# Loop thru mystery dirs
for d in *mystery*/ ; do
    echo "$d"
    echo $(ls)
    cd $d
done