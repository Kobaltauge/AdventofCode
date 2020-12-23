cups = [3,8,9,1,2,5,4,6,7]

cupslenght = len(cups)
index = 0

for move in range(10):

    destination = cups[index]-1
    idx1 = (index+1) % cupslenght
    idx2 = (index+2) % cupslenght
    idx3 = (index+3) % cupslenght
    idxpop = (index+4) % cupslenght
    pickup = [cups[idx1], cups[idx2], cups[idx3]]
    popidx = [idx1, idx2, idx3]
    for idx in popidx:
        cups[idx] = 99
    
    cups.insert(index+1,2)
    
    found = False
    cupssorted = cups.copy()
    cupssorted.sort()

    while not found:
        if destination in cups:
            destidx = cups.index(destination)
            found = True
        else:
            destination -= 1
            if destination < cupssorted[0]:
                destination = cupssorted[-3]
            found = False

    counter = 0
    for val in pickup:
        cups[popidx[counter]+1] = val
        counter += 1
    # cups[destidx+1:destidx+1] = pickup

    index += 1
    if index >= cupslenght:
        index = 0

print("part 2")