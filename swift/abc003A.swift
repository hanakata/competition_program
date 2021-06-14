let n = readLine()!
var i = 1;
var m = 0;
while i <= Int(n)!{
  m = m +  i * 10000;
  i = i + 1;
}
let s = m / Int(n)!;
print(s);