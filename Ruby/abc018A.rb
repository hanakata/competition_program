n = [0,0,0]
r = [1,1,1]
i = 0
while line = STDIN.gets
  l = line.chomp!.to_i
  a = l
  n[i] = a
  i = i + 1
end
for i in 0..2 do
  for j in 0..2 do
    if n[i] < n[j] then
      r[i] = r[i] + 1
    end
  end
end

print r[0], "\n"
print r[1], "\n"
print r[2], "\n"
