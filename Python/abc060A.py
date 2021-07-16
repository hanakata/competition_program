n = input().split()
a = list(n[0])
b = list(n[1])
c = list(n[2])
if a[-1] == b[0] and b[-1] == c[0]:
  print("YES")
else:
  print("NO")