n = list(input())
if n[0] == n[1]:
  if n[0] == n[2]:
    print("Yes")
  else:
    print("No")
elif n[1] == n[2]:
  if n[1] == n[3]:
    print("Yes")
  else:
    print("No")
else:
  print("No")
