var n = [0, 0, 0]
var r = [1, 1, 1]
var i = 0
while let a: String = readLine() {
  n[i] = Int(a)!
  i = i + 1
}
for i = 0; i < 3; i++ {
  for var j = 0; j < 3; j++ {
    if(n[i] < n[j]) {
      r[i] = r[i] + 1
    }
  }
}
print(r[0])
print(r[1])
print(r[2])
