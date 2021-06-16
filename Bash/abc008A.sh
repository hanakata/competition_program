#!/bin/bash
read a
set -- $a
echo `expr $2 - $1 + 1`
