#!/bin/bash
read a
set -- $a
if [ $1 -lt $2 ]; then
  echo "$2"
else
  echo "$1"
fi