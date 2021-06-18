a = gets.split(" ")
n = a[0].to_i
m = a[1].to_i
s = n % m
if s == 0 then
  print "YES\n"
else
  print "NO\n"
end
