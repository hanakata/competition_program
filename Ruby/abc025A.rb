a = gets.split("")
b = gets.to_i
p = (b - 1) / 5
q = b % 5
if q == 0 then
  s = 4
else
  s = q - 1
end
print a[p] + a[s],"\n"
