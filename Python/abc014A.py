n = input()
m = input()
i = int(n) % int(m)
if i == 0:
  print(0)
else:
  print(int(m) - int(i))

