
fn main() {
    let mut x: i32 = 0;
    let max_loop = 1000000000;

    for _i in 0..max_loop {
        x = x+1;
    }

    println!("x: {}", x);
}
