#!/bin/bash
read m
declare -i m
if [ $m -lt 100 ]; then
  echo "00"
fi

if [ $m -ge 100 -a $m -le 5000 ]; then
  n=`expr $m / 100`
  printf "%02d\n" "${n}"
fi

if [ $m -ge 6000 -a $m -le 30000 ]; then
  echo `expr $m / 1000 + 50`
fi

if [ $m -ge 35000 -a $m -le 70000 ]; then
  n_t=`expr $m / 1000 - 30`
  echo `expr $n_t / 5 + 80`
fi

if [ $m -gt 70000 ]; then
  echo "89"
fi