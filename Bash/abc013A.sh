#!/bin/bash
read n
str=("A" "B" "C" "D" "E")
i=0
while [ $str[$i] -ne $n ]
do
  i=`expr $i + 1`
done
echo `expr $i + 1`
