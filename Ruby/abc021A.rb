l = []
c = [8,4,2,1]
i = 0
n = gets.to_i
for a in c do
  if n == 0 then
    break
  end
  if a <= n then
    n = n - a
    l.push(a)
    i = i + 1
  end
end
print i, "\n"
for b in l do
  print b, "\n"
end