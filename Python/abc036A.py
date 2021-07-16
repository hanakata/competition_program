n = input().split()
if int(n[1]) % int(n[0]) == 0:
  s = int(n[1]) / int(n[0])
else:
  s = int(n[1]) / int(n[0]) + 1
print(int(s))