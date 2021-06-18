n = gets.to_i
m = gets.to_i
i = n % m
if i == 0 then
  print "0\n"
else
  print m - i,"\n"
end
