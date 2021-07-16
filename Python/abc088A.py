n = ['A', 'B', 'C', 'D', 'E', 'F']
m = input().split()
if n.index(m[0]) < n.index(m[1]):
  print("<")
elif n.index(m[0]) > n.index(m[1]):
  print(">")
else:
  print("=")