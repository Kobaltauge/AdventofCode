cups = [3,8,9,1,2,5,4,6,7]

cupslenght = len(cups)
index = 0

for move in range(10):
    destination = cups[index]-1
    idx1 = index+1
    if idx1 >= cupslenght:
        idx1 = index % cupslenght
    idx2 = index+2
    if idx2 >= cupslenght:
        idx2 = index % cupslenght
    idx3 = index+3
    if idx3 >= cupslenght:
        idx3 = index % cupslenght

    pickup = [cups[idx1], cups[idx2], cups[idx3]]

    found = False

    while not found:
        if destination <= 0:
            destination = 9
            found = False
        elif destination not in pickup:
            destidx = cups.index(destination)
            found = True
        else:
            destination -= 1
            found = False

    cupswork[destidx + 1:destidx + 1] = pickup
    cupswork[destidx + 1 + cupslenght:destidx + 1 + cupslenght] = pickup
    if destidx < index:
        cups = cupswork[destidx + 3:cupslenght + 3]
    else:   
        cups = cupswork[:cupslenght]
    
    index += 1
    if index >= cupslenght:
        index = 0

print("part 2")