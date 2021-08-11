read a
set -- $a
if [ $1 -eq $2 ]; then
  echo $3
  exit
fi
if [ $1 -eq $3 ]; then
  echo $2
  exit
fi
if [ $2 -eq $3 ]; then
  echo $1
  exit
fi