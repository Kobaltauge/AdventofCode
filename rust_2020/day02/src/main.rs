use std::fs;
use regex::Regex;

fn check_passwd(s: &str, re: &Regex) -> Option<String> {
    let matches = re.captures(s)?;
    println!("{}",&matches[0]);
    Some(s.to_string())
}

fn main() {
    let values = fs::read_to_string("input.txt").expect("could not load");
    let regex = Regex::new(r"^(\d+)-(\d+) (.): (.*)$").unwrap();
    values
        .split('\n')
        .filter_map(|s| check_passwd(s, &regex).Some());
}
        // .collect()
        // .to_string()
// }


