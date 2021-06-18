#!/bin/bash
read a
set -- $a
m=`expr $1 % $2`
if [ $m -eq 0 ]; then
  echo "YES"
else
  echo "NO"
fi
