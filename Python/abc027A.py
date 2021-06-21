import collections
n = input().split()
nd = collections.Counter(n)
for m, o in nd.items():
  if o == 3:
    a = m
  elif o == 1:
    a = m
print(a)