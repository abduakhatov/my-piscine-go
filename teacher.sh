#!/bin/bash
set -e
echo $(ls)
echo ">>>"
echo $(find -name "*mystery*")
echo ">>>"
for d in *mystery*/ ; 
do 
    if [ -z "$d" ]
        then
            # echo "$inter_id continue"
            continue        
    fi
    cd $d
    # GENDER and HEIGHT 
    scene=$(grep "CLUE" crimescene | head -n 1)
    gender=$( echo "$scene" | cut -d ' ' -f18 | tr -d ,)
    height=$( echo "$scene" | cut -d ' ' -f21 | tr -d \'.)
    # PEOPLE INFO
    witness=$(grep "CLUE" crimescene | tail -n 1 | grep -oP '(?<= latte was )\w.*' | cut -d ',' -f1)
    grep $witness people | while read -r people ; do
        w_full_name=$( echo "$people" | awk '{ print $1, $2 }')
        w_gender=$( echo "$people" | awk '{ print $3 }')
        w_age=$( echo "$people" | awk '{ print $4 }')
        w_adress=$( echo "$people" | awk  '{ printf("%s_%s",$5,$6) }' | tr -d ',')
        w_line=$( echo "$people" | awk '{ print $8 }')
        # find INTERVIEW by ADDRESS
        inter_id=$(head -n $w_line streets/$w_adress | tail -n 1 | cut -f 2 -d '#' )
        interview=$(cat interviews/interview-$inter_id | grep -oP '(?<= Describes it as )\w.*' | tr -d '\",')
        if [ -z "$interview" ]
        then
            # echo "$inter_id continue"
            continue        
        fi
        export interview_id=$inter_id
        echo $interview_id
        echo $(cat interviews/interview-$inter_id)
        car_color=$( echo "$interview" | cut -d ' ' -f2)
        car_model=$( echo "$interview" | cut -d ' ' -f3)
        car_plate_start=$( echo "$interview" | cut -d ' ' -f11)
        car_plate_end=$( echo "$interview" | cut -d ' ' -f15)
        # search for VEHICLE
        suspects=$(grep -A 5 -E "License Plate $car_plate_start.*9$" vehicles | grep -A 4 -B 1 'Honda' | grep -A 3 -B 2 'Blue' | grep -A 1 -B 4 'Height: 6')
        suspects=($(echo "|$suspects" | awk 'BEGIN {ORS = "|" } { print }' | tr -s ' ' '_' | tr -s '--' '-' | tr -s '-' ' '))
        main_suspect=0
        ind=1
        for suspect in "${suspects[@]}";
        do
            sus_plate=$(echo "$suspect" | cut -d '|' -f 2 | cut -d ':' -f 2 | cut -d '_' -f 3)
            sus_make=$(echo "$suspect" | cut -d '|' -f 3 | cut -d ':' -f 2 | tr -s '_' ' ' | xargs)
            sus_color=$(echo "$suspect" | cut -d '|' -f 4 | cut -d ':' -f 2 | tr -s '_' ' ' | xargs)
            sus_name=$(echo "$suspect" | cut -d '|' -f 5 | cut -d ':' -f 2 | tr -s '_' ' ' | xargs)
            sus_height=$(echo "$suspect" | cut -d '|' -f 6 | cut -d ':' -f 2 | tr -s '_' ' ' | tr -d ' ')
            sus_weight=$(echo "$suspect" | cut -d '|' -f 7 | cut -d ':' -f 2 | tr -s '_' ' ' | xargs | cut -d " " -f 1)
            sus_weight_text='female'

            echo "$sus_plate"
            echo "$sus_make"
            echo "$sus_color"
            echo "$sus_name"
            echo "$sus_height"

            if [ $sus_weight -gt 200 ]
            then 
                sus_weight_text='male'
            fi
            tmp_count=$(cat memberships/AAA memberships/Delta_SkyMiles memberships/Terminal_City_Library memberships/Museum_of_Bash_History | grep -c "${sus_name}")
            if [ $tmp_count > 4 ] && [ $sus_weight_text == $gender ]
            then
                main_suspect=$ind
            fi
            ind=$((ind+1))
        done
        echo $main_suspect
        export MAIN_SUSPECT=$main_suspect
    done
done