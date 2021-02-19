use std::fs;
use regex::Regex;

fn check_passwd(s: &str, re: &Regex) -> Option<bool> {
    let matches = re.captures(s)?;
    // println!("{}",&matches[1]);
    // println!("{}",&matches[2]);
    // println!("{}",&matches[3]);
    // println!("{}",&matches[4]);
    let min = &matches[1].parse::<usize>().unwrap();
    let max = &matches[2].parse::<usize>().unwrap();
    let count = &matches[4].matches(&matches[3]).count();
    Some(count >= min && count <= max)
}

fn main() {
    let mut count = 0;
    let values = fs::read_to_string("input.txt").expect("could not load");
    let regex = Regex::new(r"^(\d+)-(\d+) (.): (.*)$").unwrap();
    for item in values.split('\n').into_iter() {
        match check_passwd(&item, &regex) {
            Some(true) => count = count + 1,
            _ => (),
        }
    }
    println!("{}", count)
        // .collect()
        // .to_string()
        // .filter_map(|s| check_passwd(s, &regex));
}


