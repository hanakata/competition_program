#!/bin/bash
read a
if [ $a -eq 12 ]; then
  echo 1
else
  echo `expr $a + 1`
fi
