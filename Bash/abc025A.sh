#!/bin/bash
read s
read n
a=$((n/5))
p=$((n%5))
if [ $p -eq 0 ]; then
  q=4
  a=$((n/5-1))
else
  q=$((p-1))
fi
echo ${s:$a:1}${s:$q:1}
