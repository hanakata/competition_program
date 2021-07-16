n = list(map(int,input().split()))
if abs(n[0] - n[1]) < abs(n[0] - n[2]):
  print("A")
else:
  print("B")