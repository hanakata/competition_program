n = input().split()
m = list(map(int, n))
if m[1] - m[0] == m[2] - m[1]:
  print("YES")
else:
  print("NO")