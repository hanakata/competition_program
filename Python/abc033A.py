import collections
n = list(input())
nd = collections.Counter(n)
for m, o in nd.items():
  if o == 4:
    print("SAME")
  else:
    print("DIFFERENT")
    break
