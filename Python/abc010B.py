n = int(input())
m = input().split()
s = 0
for l in m:
    if int(l) == 9:
        s += 0
    elif int(l) >= 7:
        s += int(l) - 7
    elif int(l) >= 3:
        s += int(l) - 3
    elif int(l) >= 1:
        s += int(l) - 1
print(s)
