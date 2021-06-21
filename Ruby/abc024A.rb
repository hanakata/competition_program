p = gets.split(" ")
q = gets.split(" ")
n = q[0].to_i * p[0].to_i
m = q[1].to_i * p[1].to_i
o = q[0].to_i + q[1].to_i
if p[3].to_i <= o then
  print n + m - o * p[2].to_i,"\n";
else
  print n + m,"\n";
end
