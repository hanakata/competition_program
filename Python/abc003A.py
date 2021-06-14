m = input()
i = int(m)
j = 1
b = 0
 
while j <= i:
    a = j * 10000
    b += a
    j += 1
 
s = b / i
print(int(s))