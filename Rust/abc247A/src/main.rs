fn solved(m: String) -> String {
    "0".to_owned() + &m[..3]
}

fn main() {
    let mut m = String::new();
    std::io::stdin().read_line(&mut m);
    println!("{}", solved(m));
}