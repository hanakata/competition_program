let n = readLine()!
let m = readLine()!
let i = Int(n)! % Int(m)!
if(Int(i) == 0 ){
  print(0)
}else{
  print(Int(m) - Int(i))
}
