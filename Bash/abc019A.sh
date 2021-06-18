#!/bin/bash
m=0
declare -i a b c
read a b c
n=($a $b $c)
for ((i=0; i < 2; i++)); do
  for ((j=2; j > i; j--)); do
    m=`expr $j - 1`
    if [ ${n[j]} -lt ${n[m]} ]; then
      t=${n[m]}
      n[$m]=${n[j]}
      n[$j]=$t
    fi
  done
done
echo ${n[1]}
