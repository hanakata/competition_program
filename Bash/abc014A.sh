#!/bin/bash
read n
read m
declare -i n
declare -i m
i=`expr $n % $m`
if [ $i -eq 0 ]; then
  echo "0"
else
  echo `expr $m - $i`
fi
