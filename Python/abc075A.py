n = list(map(int,input().split()))
n.sort()
if n.count(n[0]) == 2:
  print(n[2])
else:
  print(n[0])