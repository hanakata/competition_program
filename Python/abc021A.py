l = []
c = [8,4,2,1]
i = 0
n = int(input())
for a in c:
    if(n == 0):
      break
    if(a <= n):
      n = n - a
      l.append(a)
      i = i + 1
print(i)
for b in l:
    print(b)
