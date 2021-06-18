let a = readLine()!
let b = a.characters.split{$0 == " "}.map(String.init)
let m = Int(b[0])! % Int(b[1])!
if ( m == 0 ) {
    print("YES")
} else {
    print("NO")
}
