
// fn main() {

//     let mut x :i64 = 0;
//    while x < 1000000000 {
//       x+=1;     
//    }
//    println!("outside loop x value is {}",x);
// }

//optimised code - for loop instead of a while loop - ChatGPT!
fn main(){
   let mut x:i64 = 0;
    for i in 0..1000000000 {
        x += i;
    }    
    println!("x: {}", x);
}

