use std::{thread, time};

fn main() {
    let mut v: [u32; 1000000] = [0; 1000000];
    
    // create a concatenate list where the value is
    // simultaneous the index of the next element

    let mut index = 0;
    for i in &[3,8,9,1,2,5,4,6,7] {
        v[index] = *i as u32;
        index = *i as usize;
    }
    
    for i in 10..1000001 {
        v[index] = i as u32;
        index = i as usize;
    }
    
    
    thread::sleep(time::Duration::from_secs(100));
    // let mut index = 0;
    // let mut moveno = 1;
    // // let moves = 10;
    // let moves = 10000000;
    // let vlen = v.len();
    // for _ in 0..moves {
    //     // println!("\n-- move {} --", moveno);
    //     // println!("cups {:?}", v);
    //     // println!("index {:?}", index);

    //     let idx1 =
    //         if index + 1 >= vlen {
    //             index + 1 - vlen
    //         } else {
    //             index + 1
    //         };
    //     let idx2 =
    //         if index + 2 >= vlen {
    //             index + 2 - vlen
    //         } else {
    //             index + 2
    //         };
    //     let idx3 =
    //         if index + 3 >= vlen {
    //             index + 3 - vlen
    //         } else {
    //             index + 3
    //         };
    //     let mut pickup: Vec<i32> = Vec::new();
    //     pickup.push(v[idx1]);
    //     pickup.push(v[idx2]);
    //     pickup.push(v[idx3]);
    //     // println!("pickup {:?}", pickup);
        
    //     let destination = v[index] as i32;
    //     let mut target = if destination > 1 {destination - 1} else {maxno};
    //     // println!("destination {:?}", target);
    //     while pickup[0] == target || pickup[1] == target || pickup[2] == target {
    //         target = if target > 1 {target - 1} else {maxno};
    //     }
    //     // println!("destinationafter {:?}", target);
        
    //     v.retain(|&x| ! pickup.contains(&x));
    //     let destidx = getindex(target, &v);
    //     // println!("destinationindex {:?}", destidx);

    //     v.insert(destidx + 1, pickup[0]);
    //     v.insert(destidx + 2, pickup[1]);
    //     v.insert(destidx + 3, pickup[2]);

    //     while v[index] != destination {
    //         // v.rotate_left(1);
    //         index = if index + 1 <= moves {0} else {index +1};
    //     }

    //     // thread::sleep(time::Duration::from_secs(1));

    //     index = if index + 1 <= moves {0} else {index +1};
    //     if moveno % 100 == 0 {println!("{}",moveno)};
    //     moveno = moveno + 1;
    // }
    // println!("{:?}", v);
    // println!("{}", v[getindex(1, &v)+1]);
    // println!("{}", v[getindex(1, &v)+2]);
    // println!("{}", v[getindex(1, &v)+1] * v[getindex(1, &v)+2]);
}
