import math
 
m = input()
i = int(m)
 
if i < 100:
    print("00")
if 100 <= i <= 5000:
    n_t = i / 1000 * 10
    n = int(n_t)
    check = int(math.log10(n))
    if check == 0:
        print("0"+str(n))
    else:
        print(n)