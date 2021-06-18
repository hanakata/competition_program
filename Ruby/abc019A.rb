line = gets.split(" ")
l = line.map(&:to_i)
l.sort!
print l[1],"\n"