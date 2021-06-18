import fileinput
n = [0,0,0]
r = [1,1,1]
i = 0
for line in fileinput.input():
    l = line
    n[i] = int(l)
    i += 1
for i in range(3):
  for j in range(3):
    if(n[i] < n[j]):
      r[i] += 1

print(r[0])
print(r[1])
print(r[2])
