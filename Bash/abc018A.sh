#!/bin/bash
n=()
declare -a r=("1" "1" "1")

i=0
declare -i a
while read a; do
  n[$i]=$a
  i=`expr $i + 1`
done
for((i=0; i < 3; i++)); do
  for((j=0; j < 3; j++)); do
    if [ ${n[i]} -lt ${n[j]} ]; then
      r[$i]=`expr ${r[i]} + 1`
    fi
  done
done

echo ${r[0]}
echo ${r[1]}
echo ${r[2]}
