cups = [3,8,9,1,2,5,4,6,7]

cupslenght = len(cups)
index = 0

for move in range(10):
    cupswork = cups + cups
    destination = cupswork[index]-1
    idx1 = index+1
    idx2 = index+2
    idx3 = index+3
    pickup = [cupswork[idx1], cupswork[idx2], cupswork[idx3]]
    cupswork = cupswork[:idx1] + cupswork[idx3+1:cupslenght] + \
               cupswork[cupslenght:cupslenght + idx1] + \
               cupswork[cupslenght + idx3+1:] 

    found = False
    cupssorted = cupswork.copy()
    cupssorted.sort()

    while not found:
        if destination in cupswork:
            destidx = cupswork.index(destination)
            found = True
        elif destination < cupssorted[0]:
            destination = cupssorted[-1]
            found = False
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