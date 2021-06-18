#!/bin/bash
read l
read s
if [ ${#l} -lt ${#s} ]; then
  echo $s
else
  echo $l
fi
