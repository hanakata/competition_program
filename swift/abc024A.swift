let a = readLine()!
let b = readLine()!
let p = a.characters.split{$0 == " "}.map(String.init)
let q = b.characters.split{$0 == " "}.map(String.init)
let n = Int(q[0])! * Int(p[0])!
let m = Int(q[1])! * Int(p[1])!
let o = Int(q[0])! + Int(q[1])!
if(Int(p[3])! <= o){
  print(n + m - o * Int(p[2])!)
}else{
  print(n + m)
}
