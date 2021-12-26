with open("day_25.txt") as f:
    map = f.read().splitlines()
    map = [list(i) for i in map]

def printmap(east, south):
    for i in range(0,bb+1):
        row = ""
        for j in range(0,br+1):
            if [j,i] in east:
                row = row + ">"
            elif [j,i] in south:
                row = row + "v"
            else:
                row = row + "."
        # print(f"{row}")

bb = len(map)-1
br = len(map[0])-1

east = []
south = []

for i in enumerate(map):
    for j in enumerate(map[i[0]]):
        if map[i[0]][j[0]] == ">":
            east.append([j[0],i[0]])
        elif map[i[0]][j[0]] == "v":
            south.append([j[0],i[0]])

round = 0
move = 1
complete = east + south

# printmap(east, south)

while move == 1:
    move = 0
    eastnew = east.copy()
    for e in east:
        next = e[0] + 1
        if next > br:
            next = 0
        if not [next, e[1]] in complete:
            eastnew.remove(e)
            eastnew.append([next, e[1]])
            move = 1
    east = eastnew.copy()
    complete = east + south
    southnew = south.copy()
    for s in south:
        next = s[1] + 1
        if next > bb:
            next = 0
        if not [s[0], next] in complete:
            southnew.remove(s)
            southnew.append([s[0], next])
            move = 1
    south = southnew.copy()
    complete = east + south
    # print(f"\n")
    # printmap(east, south)
    round += 1


print(f"result part 1: {round}")