n = input().split()
t = int(n[1]) / int(n[0])
a = int(n[3]) / int(n[2])
if t < a:
  print("AOKI")
if t > a:
  print("TAKAHASHI")
if t == a:
  print("DRAW")