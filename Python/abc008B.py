import collections
n = int(input())
a = []

for s in range(n):
  a.append(input())

count_dict = collections.Counter(a)
for k, v in count_dict.most_common(1):
    print(k)
