use std::fs;
use regex::Regex;

fn check_passwd(s: &str, re: &Regex) -> String {
    let matches = re.captures(s)?;
    println!("{}",matches);
    matches
}

pub fn day02a() -> String {
    let values = fs::read_to_string("input_02.txt").expect("could not load");
    let regex = Regex::new(r"^(\d+)-(\d+) (.): (.*)$").unwrap();
    values
        .split('\n')
        .filter_map(|s| check_passwd(s, &regex))
        .collect()
        .to_string()
}


