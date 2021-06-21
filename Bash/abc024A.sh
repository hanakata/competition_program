#!/bin/bash

read p
read q
set -- $p
a=$1
b=$2
c=$3
k=$4
set -- $q
s=$1
t=$2
n=$(($s*$a))
m=$(($t*$b))
o=$(($s+$t))
if [ $k -le $o ] ; then
  echo $(($n+$m-$o*$c))
else
  echo $(($n+$m))
fi
