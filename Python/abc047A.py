n = input().split()
m = list(map(int, n))
m.sort()
if m[0] + m[1] == m[2]:
  print("Yes")
else:
  print("No")