a = ['a','t','c','o','d','e','r']
s = input().strip()
t = input().strip()
i = 0
j = 0
if '@' not in s and '@' not in t:
  if s == t:
    print("You can win")
    j = 1
  else:
    print("You will lose")
    j = 1
else:
  s_l = list(s)
  t_l = list(t)
  for n in s_l:
    if n == '@':
      if t_l[i] not in a and t_l[i] != '@':
        print("You will lose")
        j = 1
        break
    elif t_l[i] == '@' and n not in a:
      print("You will lose")
      j = 1
      break
    elif t_l[i] != '@' and n != t_l[i]:
      print("You will lose")
      j = 1
      break
    i = i + 1
if j == 0:
    print("You can win")