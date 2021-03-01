use std::{collections::BTreeSet, iter::FromIterator};
// use modulo::Mod;

fn pickupdel<T>(mut items: Vec<T>, to_remove: Vec<T>) -> Vec<T>
where
    T: std::cmp::Ord,
{
    let to_remove = BTreeSet::from_iter(to_remove);
    items.retain(|e| !to_remove.contains(e));
    items
}

fn destidxsearch(v: Vec<i32>, pickup: Vec<i32>, mut destination: i32) -> (i32, usize) {
    // let destination = destination;
    let destindex = 0;
    if !pickup.contains(&destination) {
        (destination, v.iter().position(|&r| r == destination).unwrap());
    } else if destination <= 0 {
        destination = 1000000;
        let (destination, destindex) = destidxsearch(v.clone(), pickup.clone(), destination.clone());
    } else {
        destination = destination - 1;
        let (destination, destindex) = destidxsearch(v.clone(), pickup.clone(), destination.clone());
    }
    return (destination, destindex)
}


fn main() {
    let mut v = vec![3,8,9,1,2,5,4,6,7];
    // let mut v = vec![3,8,9,1,2,5,4,6,7,10,11,12,13,14,15];
    // let mut v = vec![4,5,9,6,7,2,8,1,3];
    
    // for n in 10..1000001 {
    //     // println!("{}",n);
    //     v.push(n);
    // }
    
    let mut index = 0;
    let mut moveno = 1;

    // while moveno < 10000000 {
    while moveno < 10 {
        let mut vec_mut = &v;
        println!("in loop {:?}\n", v);

        let idx1 =
            if index + 1 > 1000000 {
                index % 1000000
            } else {
                index + 1
            };
        let idx2 =
            if index + 2 > 1000000 {
                index % 1000000
            } else {
                index + 2
            };
        let idx3 =
            if index + 3 > 1000000 {
                index % 1000000
            } else {
                index + 3
            };
        let mut pickup: Vec<i32> = Vec::new();
        println!("indexe {} {} {}", idx1, idx2, idx3);
        println!("pickup {:?}", pickup);
        pickup.push(vec_mut[idx1]);
        pickup.push(vec_mut[idx2]);
        pickup.push(vec_mut[idx3]);
        println!("pickup after{:?}", pickup);

        let destination = &vec_mut[index] - 1;
        println!("destination {:?}", destination);
        
        vec_mut = pickupdel(vec_mut.clone(), pickup.clone());
        // v.retain(|x| pickup.contains(x));
        // v.retain(|&x| x == pickup.iter().unwrap());
        // pickup.iter().map(|x| x % 2 == 0);

        // println!("{:?} {:?} {:?}", idx1, idx2, idx3);
        // println!("{:?}", pickup);
        // println!("{:?}", v);

        let (destination, destidx) = destidxsearch(vec_mut.clone(), pickup.clone(), destination.clone());

        println!("move {}", moveno);
        println!("new vec {:?}", vec_mut);
        println!("destination {:?}", destination);

        vec_mut.insert(destidx + 2, pickup[0]);
        vec_mut.insert(destidx + 3, pickup[1]);
        vec_mut.insert(destidx + 4, pickup[2]);

        *v = vec_mut;
        index = index + 1;
        moveno = moveno + 1;
    }
}
