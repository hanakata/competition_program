#!/bin/bash
read n
m=`expr $n % $1`
if [ $m -eq 0 ]; then
  echo "YES"
else
  echo "NO"
fi