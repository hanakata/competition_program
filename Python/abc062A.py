a = ['1','3','5','7','8','10','12']
b = ['4','6','9','11']
n = input().split()
if n[0] in a and n[1] in a or n[0] in b and n[1] in b:
  print("Yes")
else:
  print("No")