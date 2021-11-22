read m
declare -i m
if [ $m -le 59 ]; then
    echo "Bad"
fi

if [ $m -ge 60 -a $m -le 89 ]; then
  echo "Good"
fi

if [ $m -ge 90 -a $m -le 99 ]; then
  echo "Great"
fi

if [ $m -eq 100 ]; then
  echo "Perfect"
fi