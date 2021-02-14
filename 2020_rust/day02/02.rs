use std::fs;

fn main() {
    let inputfile = "example.txt";
    // let inputfile = "input.txt";

    let values = fs::read_to_string(inputfile).expect("could not load");
    let values = values
        .split('\n')
        .filter_map(|s| s.parse::<u32>().ok())
        .collect::<Vec<u32>>();
    
    for x in &values {
        for y in &values {
            if x+y == 2020 {
                // println!("{} and {} = 2020", x,y);
                println!("{}", x * y);
                break
            }
        }
    }
    for k in &values {
        for x in &values {
            for y in &values {
                if k+x+y == 2020 {
                    // println!("{} and {} and {} = 2020", k,x,y);
                    println!("{}", k * x * y);
                    break
                }
            }
        }
    }
}
