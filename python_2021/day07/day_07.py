with open("day_07.txt") as f:
    crabs = list(map(int, (f.read().split(','))))

# crabs = [16,1,2,0,4,2,7,1,2,14]

bestpos = 0
for pos in range(0, max(crabs)+1):
    checkpos = sum(list(map(abs, [x-y for (x, y) in zip([pos] * len(crabs), crabs)])) )
    if checkpos < bestpos or bestpos == 0:
        bestpos = checkpos 
print(f"solution part 1: {bestpos}\n")

bestpos = 0
for pos in range(0, max(crabs)+1):
    checkpos = list(map(abs, [x-y for (x, y) in zip([pos] * len(crabs), crabs)]))
    checkpos = sum([sum([i for i in range(0, j+1)]) for j in checkpos])
    if checkpos < bestpos or bestpos == 0:
        bestpos = checkpos 
print(f"solution part 2: {bestpos}\n")
