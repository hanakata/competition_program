n = list(map(int,input().split()))
m = n[2] - n[1]
if n[2] <= n[1]:
  print("delicious")
elif m <= n[0]:
  print("safe")
else:
  print("dangerous")