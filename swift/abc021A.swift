var l: [Int] = []
var c = [8, 4, 2, 1]
let a = readLine()!
var n = Int(a)!
var i = 0
for a in c {
  if (n == 0) {
    break
  }
  if (a <= n){
    n = n - a
    l.append(a)
    i = i + 1
  }
}
print (i)
for b in l {
  print(b)
}
