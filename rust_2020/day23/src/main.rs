// use std::{thread, time};
use std::{thread};

const STACK_SIZE: usize = 4 * 1024 * 1024;

fn run() {

    let mut v: [u32; 1000001] = [0; 1000001];

    // create a concatenate list where the value is
    // simultaneous the index of the next element
    // let startv = [3,8,9,1,2,5,4,6,7];
    let startv = [4,5,9,6,7,2,8,1,3];

    for i in 0..8 {
        v[startv[i]] = startv[i+1] as u32;
    }

    v[startv[startv.len()-1]] = 10;
    for i in 10..1000000 {
        v[i] = i as u32 + 1;
    }
    
    // close the loop
    v[1000000] = startv[0] as u32;
    
    // thread::sleep(time::Duration::from_secs(100));

    let mut currcup = startv[0] as u32;
    let moves = 10000000;
    let maxno = 1000000;

    let mut moveno = 1;
    for _ in 0..moves {
        let pickup1 = v[currcup as usize];
        let pickup2 = v[pickup1 as usize];
        let pickup3 = v[pickup2 as usize];
        let cupafterpickup = v[pickup3 as usize];

        let mut dest = if currcup > 1 {currcup - 1} else {maxno};
        while pickup1 == dest || pickup2 == dest || pickup3 == dest {
            dest = if dest > 1 {dest - 1} else {maxno};
        }
        
        // in the concatenated list this is before the shift
        // 
        // currcup => pickup1 => pickup2 => pickup3 => cupafterpickup
        //
        // after that the pickups are shiftet after the destiantion cup
        // therefore we have two separte concatenation
        //
        // dest => pickup1 => pickup2 => pickup3 => cupafterdest
        // currcup => cupafterpickup
        
        let cupafterdest = v[dest as usize];
        v[pickup3 as usize] = cupafterdest;
        v[dest as usize] = pickup1;
        v[currcup as usize] = cupafterpickup;

        // println!("\n-- move {} --", moveno);
        // println!("{:?}", v);
        // println!("index {:?}", currcup);
        // println!("pickup {:?} {:?} {:?}", pickup1, pickup2, pickup3);
        // println!("destination {:?}", dest);
        // thread::sleep(time::Duration::from_secs(1));
        
        currcup = cupafterpickup;
        moveno = moveno + 1;
    }
    
    // println!("{:?}", v);
    println!("{} {}", v[1], v[(v[1] as usize)]);
    println!("{}", v[1] as u64 * v[(v[1] as usize)] as u64);
    // println!("ENDE");
}

fn main() {
    let child = thread::Builder::new()
        .stack_size(STACK_SIZE)
        .spawn(run)
        .unwrap();

    child.join().unwrap();
}
