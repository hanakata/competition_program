n = input().split()
m = list(map(int,n))
a = m[0] * m[1]
b = m[2] * m[3]
if a < b:
  print(b)
else:
  print(a)