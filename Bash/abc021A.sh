#!/bin/bash
l=()
declare -a c=("8" "4" "2" "1")
i=0
j=0
declare -i n
read n
while [ $n -ne 0 ]; do
  b=${c[j]}
  if [ $b -le $n ]; then
    n=`expr $n - $b`
    l[$i]=$b
    i=`expr $i + 1`
  fi
  j=`expr $j + 1`
done
j=0
echo $i
while [ $j -lt $i ]; do
  echo ${l[j]}
   j=`expr $j + 1`
done
