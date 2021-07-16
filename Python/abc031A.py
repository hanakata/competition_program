n = input().split()
a = (int(n[0]) + 1) * int(n[1])
d = (int(n[1]) + 1) * int(n[0])
if a < d:
  print(d)
else:
  print(a)