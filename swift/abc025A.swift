let a = readLine()!
let b = readLine()!
let a_list = a.characters.map { String($0) }
let p = (Int(b)! - 1) / 5
let q = Int(b)! % 5
var s = 0
if(q == 0){
  s = 4
}else{
  s = q - 1
}
print(a_list[p] + a_list[s])
