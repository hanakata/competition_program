n = int(input())
a = []

for s in range(n):
  a.append(int(input()))

u = list(set(a))
u.sort()
print(u[n - 2])