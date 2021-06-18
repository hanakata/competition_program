var m = 0
while let a: String = readLine() {
  var b = a.characters.split{$0 == " "}.map(String.init)
  var n = Int(b[0])! * Int(b[1])! / 10
  m = m + n
}
print (m)
