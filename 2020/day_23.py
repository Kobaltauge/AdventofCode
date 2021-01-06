cups = [3,8,9,1,2,5,4,6,7]
cups = [4,5,9,6,7,2,8,1,3]

cupslenght = len(cups)
index = 0

for move in range(100):
    cupswork = cups.copy()
    cupswork = cupswork[index:] + cupswork[:index] 
    destination = cupswork[0] - 1
    pickup = [cupswork[1], cupswork[2], cupswork[3]]
    cupswork = cupswork[:1] + cupswork[4:] 

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

    cups = cupswork[-index:] + cupswork[:-index] 
    
    index += 1
    if index >= cupslenght:
        index = 0

erg = cups[cups.index(1)+1:] + cups[:cups.index(1)]
print(erg)
print("part 2")



cups = [3,8,9,1,2,5,4,6,7]
# cups = [4,5,9,6,7,2,8,1,3]

for i in range(10, 1000001):
    cups.append(i)
cupslenght = len(cups)
index = 0

for move in range(10000000):
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
    destination = cupswork[0] - 1
    pickup = [cupswork[idx1], cupswork[idx2], cupswork[idx3]]

    found = False

    while not found:
        if destination <= 0:
            destination = 1000000
            found = False
        elif destination not in pickup:
            destidx = cupswork.index(destination)
            found = True
        else:
            destination -= 1
            found = False

    cupswork[destidx + 1:destidx + 1] = pickup

    cups = cupswork[-index:] + cupswork[:-index] 

    if move % 1000 == 0:
        print(move)
    index += 1
    if index >= cupslenght:
        index = 0

erg1 = cups[cups.index(1)+1]
erg2 = cups[cups.index(1)+2]
print(erg)
