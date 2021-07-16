n = input().split()
m = list(map(int, n))
i = m[0] + m[1]
if 24 <= i:
  print(i - 24)
else:
  print(i)