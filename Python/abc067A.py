n = list(map(int,input().split()))
if n[0] % 3 == 0 or n[1] % 3 == 0 or (n[0] + n[1]) % 3 == 0:
  print("Possible")
else:
  print("Impossible")