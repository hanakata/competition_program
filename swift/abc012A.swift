let a = readLine()!
let b = a.characters.split{$0 == " "}.map(String.init)
print(b[1] + " " + b[0])
