#!/bin/bash
m=0
declare -i a b
while read a b; do
  n=`expr $a \* $b / 10`
  m=`expr $m + $n`
done
echo $m
