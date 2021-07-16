l = [ 2,3,4,5,6,7,8,9,10,11,12,13,1 ]
n = input().split()
m = list(map(int,n))
a = l.index(m[0])
b = l.index(m[1])
if int(a) < int(b):
  print("Bob")
elif int(a) > int(b):
  print("Alice")
else:
  print("Draw")
