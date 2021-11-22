#!/bin/bash
 
read Deg Dis
 
Dis2=$(( (Dis * 100 / 10 + 30) / 60 ))
 
if [ $Dis2 -ge 0 -a $Dis2 -le 2 ]; then
  W=0
fi
if [ $Dis2 -ge 3 -a $Dis2 -le 15 ]; then
  W=1
fi
if [ $Dis2 -ge 16 -a $Dis2 -le 33 ]; then
  W=2
fi
if [ $Dis2 -ge 34 -a $Dis2 -le 54 ]; then
  W=3
fi
if [ $Dis2 -ge 55 -a $Dis2 -le 79 ]; then
  W=4
fi
if [ $Dis2 -ge 80 -a $Dis2 -le 107 ]; then
  W=5
fi
if [ $Dis2 -ge 108 -a $Dis2 -le 138 ]; then
  W=6
fi
if [ $Dis2 -ge 139 -a $Dis2 -le 171 ]; then
  W=7
fi
if [ $Dis2 -ge 172 -a $Dis2 -le 207 ]; then
  W=8
fi
if [ $Dis2 -ge 208 -a $Dis2 -le 244 ]; then
  W=9
fi
if [ $Dis2 -ge 245 -a $Dis2 -le 284 ]; then
  W=10
fi
if [ $Dis2 -ge 285 -a $Dis2 -le 326 ]; then
  W=11
fi
if [ $Dis2 -ge 327 ]; then
  W=12
fi
 
if [ $Deg -le 112 -o $Deg -ge 3488 ]; then
  Dir="N"
fi
 
if [ $Deg -ge 113 -a $Deg -le 337 ]; then
  Dir="NNE"
fi
if [ $Deg -ge 338 -a $Deg -le 562 ]; then
  Dir="NE"
fi
if [ $Deg -ge 563 -a $Deg -le 787 ]; then
  Dir="ENE"
fi
if [ $Deg -ge 788 -a $Deg -le 1012 ]; then
  Dir="E"
fi
if [ $Deg -ge 1013 -a $Deg -le 1237 ]; then
  Dir="ESE"
fi
if [ $Deg -ge 1238 -a $Deg -le 1462 ]; then
  Dir="SE"
fi
if [ $Deg -ge 1463 -a $Deg -le 1687 ]; then
  Dir="SSE"
fi
if [ $Deg -ge 1688 -a $Deg -le 1912 ]; then
  Dir="S"
fi
if [ $Deg -ge 1913 -a $Deg -le 2137 ]; then
  Dir="SSW"
fi
if [ $Deg -ge 2138 -a $Deg -le 2362 ]; then
  Dir="SW"
fi
if [ $Deg -ge 2363 -a $Deg -le 2587 ]; then
  Dir="WSW"
fi
if [ $Deg -ge 2588 -a $Deg -le 2812 ]; then
  Dir="W"
fi
if [ $Deg -ge 2813 -a $Deg -le 3037 ]; then
  Dir="WNW"
fi
if [ $Deg -ge 3038 -a $Deg -le 3262 ]; then
  Dir="NW"
fi
if [ $Deg -ge 3263 -a $Deg -le 3487 ]; then
  Dir="NNW"
fi
 
if [ $W -eq 0 ]; then
  Dir="C"
fi
 
echo "$Dir $W"