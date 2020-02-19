cd the-final-cl-test/mystery

# GENDER and HEIGHT 
scene=$(grep "CLUE" crimescene | head -n 1)
gender=$( echo "$scene" | cut -d ' ' -f18 | tr -d ,)
height=$( echo "$scene" | cut -d ' ' -f21 | tr -d \'.)
# echo "GENDER: $gender"
# echo "HEIGHT: $height"
memberships=$(grep "CLUE" crimescene | grep -oP '(?<=and membership cards for )\w.*' | cut -d '.' -f 1)
memberships=$( echo "${memberships//the /}" )
memberships=$( echo "${memberships//and /}" )
echo $memberships
# MEMBERSHIP


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
    suspects=($(grep -A 5 -e "L337.*9" vehicles | grep -B 1 -A 4 "Make: Honda" | grep -B 2 -A 3 "Color: Blue" | awk 'BEGIN {ORS = "@" } { print }' | tr -s ' ', '_' | tr -s "--" , '-' | tr -s '-' ' '))
    for suspect in "${suspects[@]}";
    do 
        # echo $suspect
        plate=$( echo $suspect | cut -d ':' -f 2)
        echo $plate
    done


done

