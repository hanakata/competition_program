n = gets.to_i
i = 1
m = 0
while i <= n do
  m +=  i * 10000
  i += 1
end
s = m / n
print s, "\n"