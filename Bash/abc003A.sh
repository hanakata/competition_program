#!/bin/bash
read n
declare -i n
i=1
m=0
while [ $i -le $n ]
do
  m=`expr $m + $i \* 10000`
  i=`expr $i + 1`
done
s=`expr $m / $n`
echo $s