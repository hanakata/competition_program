n = input().split()
if int(n[0]) < int(n[1]):
  m = int(n[2]) / int(n[0])
else:
  m = int(n[2]) / int(n[1])
print(int(m))