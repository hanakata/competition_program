n = list(map(int,input().split()))
if n[0] + n[1] > n[2] + n[3]:
  print("Left")
elif n[0] + n[1] < n[2] + n[3]:
  print("Right")
else:
  print("Balanced")