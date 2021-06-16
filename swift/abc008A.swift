let a = readLine()!
let b = a.characters.split{$0 == " "}.map(String.init)
let n = Int(b[0])!
let m = Int(b[1])!
print(m - n + 1)
