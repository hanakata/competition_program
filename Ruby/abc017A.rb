m = 0
while line = STDIN.gets
  l = line.chomp!.split.map!(&:to_i)
  a, b = l[0], l[1]
  n =  a * b / 10
  m = m + n
end
print m, "\n"
