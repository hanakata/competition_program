fn solved(m: i32, n: i32) -> i32 {
    m - n
}

fn main() {
    let mut m = String::new();
    let mut n = String::new();
    std::io::stdin().read_line(&mut m);
    std::io::stdin().read_line(&mut n);
    let m: i32 = match m.trim().parse() {
        Ok(num) => num,
        Err(_) => return,
    };
    let n: i32 = match n.trim().parse() {
        Ok(num) => num,
        Err(_) => return,
    };
    println!("{}", solved(m,n));
}

