n = list(map(int,input().split()))
if n[0] * n[1] <= n[2]:
  print(n[0] * n[1])
else:
  print(n[2])