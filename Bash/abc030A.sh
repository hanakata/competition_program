read a b c d
declare -i a
declare -i b
declare -i c
declare -i d
t=`bc <<< "scale=3; $b/$a"`
s=`bc <<< "scale=3; $d/$c"`
result=`echo "$t > $s" | bc`
if [ $result -eq 1 ]; then
    echo "TAKAHASHI"
fi

result=`echo "$t < $s" | bc`
if [ $result -eq 1 ]; then
    echo "AOKI"
fi

result=`echo "$t == $s" | bc`
if [ $result -eq 1 ]; then
    echo "DRAW"
fi