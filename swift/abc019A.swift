let a = readLine()!
var b = a.characters.split{$0 == " "}.map(String.init)
for var i = 0; i < 2; i++ {
  for var j = 2; j > i; j-- {
    if(Int(b[j])! < Int(b[j-1])!) {
      var t = b[j-1]
      b[j-1] = b[j]
      b[j] = t
    }
  }
}
print (Int(b[1])!)