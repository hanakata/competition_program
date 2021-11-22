fn main() {
    let mut a = String::new();
    std::io::stdin().read_line(&mut a);
    let v: Vec<&str> = a.split(' ').collect();
    let m: i32 = match v[0].trim().parse() {
        Ok(num) => num,
        Err(_) => return,
    };
    let n: i32 = match v[1].trim().parse() {
        Ok(num) => num,
        Err(_) => return,
    };

    let ans =
        if n < m {
            m
        } else {
            n
        };
    println!("{}", ans);
}