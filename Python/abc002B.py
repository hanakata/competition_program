d = ["a","i","u","e","o"]
n = list(input())
b = ""
for i in n:
  if i not in d:
    b += i
print(b)