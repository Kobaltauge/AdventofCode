import copy

with open("day_17_exampl.txt") as f:
    cubes = list(f.read().split())

# with open("day_17.txt") as f:
#     cubes = list(f.read().split())

cubes = [[char for char in r] for r in cubes]

pocketdim = {}

for y in range(len(cubes)):
    for x in range(len(cubes[y])):
        pocketdim[x,y,0] = cubes[x][y]


def check_adjacent(pocketdim, cube):
    adja = []
    cubex, cubey, cubez = cube
    listtmp = [x[0] for x in list(pocketdim.keys())]
    minx = listtmp.index(min(listtmp))
    maxx = listtmp.index(max(listtmp))
    listtmp = [x[1] for x in list(pocketdim.keys())]
    miny = listtmp.index(min(listtmp))
    maxy = listtmp.index(max(listtmp))
    listtmp = [x[2] for x in list(pocketdim.keys())]
    minz = listtmp.index(min(listtmp))
    maxz = listtmp.index(max(listtmp))
 
    if (miny <= cubey-1 <= maxy) and (minx <= cubex-1 <= maxx):
        adja.append(pocketdim[cubey-1][cubex-1])
    if (miny <= cubey-1 <= maxy) and (minx <= cubex <= maxx):
        adja.append(pocketdim[cubey-1][cubex])
    if (miny <= cubey-1 <= maxy) and (minx <= cubex+1 <= maxx):
        adja.append(pocketdim[cubey-1][cubex+1])
    if (miny <= cubey <= maxy) and (minx <= cubex-1 <= maxx):
        adja.append(pocketdim[cubey][cubex-1])
    if (miny <= cubey <= maxy) and (minx <= cubex+1 <= maxx):
        adja.append(pocketdim[cubey][cubex+1])
    if (miny <= cubey+1 <= maxy) and (minx <= cubex-1 <= maxx):
        adja.append(pocketdim[cubey+1][cubex-1])
    if (miny <= cubey+1 <= maxy) and (minx <= cubex <= maxx):
        adja.append(pocketdim[cubey+1][cubex])
    if (miny <= cubey+1 <= maxy) and (minx <= cubex+1 <= maxx):
        adja.append(pocketdim[cubey+1][cubex+1])
    # print(pocketdim[cubey][cubex])
    seat = pocketdim[cubey][cubex]
    if seat == ".":
        return "."
    elif (seat == "L") and (adja.count("#") == 0):
        return "#"
    elif (seat == "#") and (adja.count("#") >= 4):
        return "L"
    else:
        return seat
    
counter = 0

while True:
    newcubes = copy.deepcopy(cubes)
    for y in range(len(cubes)):
        for x in range(len(cubes[y])):
            newcubes[y][x] = check_adjacent(cubes, y, x)
    # print("\n")
    # [print(r) for r in cubes]
    # print("\n")
    # [print(r) for r in newcubes]
    counter += 1
    if not (newcubes == cubes):
        cubes = copy.deepcopy(newcubes)
#        [print(r) for r in cubes]
#        print("\n")
    else:
        break
print(sum([s.count("#") for s in cubes]))
print("part 2")