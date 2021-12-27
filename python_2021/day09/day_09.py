import math

with open("day_09.txt") as f:
    input = f.read()
    map = input.splitlines()
    maplist = [int(i) for i in input.replace('\n','')]

dimy, dimx = len(map), len(map[0])
max = dimy*dimx

def get_neigh(x, y,map):
    dimy, dimx = len(map), len(map[0])
    dir = {'up':(0,-1),'down':(0,1),'left':(-1,0),'right':(1,0)}
    if y == 0:
        del dir['up']
    if y >= dimy-1:
        del dir['down']
    if x == 0:
        del dir['left']
    if x >= dimx-1:
        del dir['right']
    return [int(map[(y + dir[i][1])][(x + dir[i][0])]) for i in dir]

minis = []
for y in range(0, dimy):
    for x in range(0, dimx):
        neigh = get_neigh(x,y, map)
        minimum = min(neigh)
        if int(map[y][x]) < minimum:
            minis.append([(x,y), int(map[y][x])])

print(f"result part 1: {sum([i[1]+1 for i in minis])}")

def get_adjacent_cells(grid, cell):
    # adjdict = {'n':(0,-1),'ne':(1,-1),'e':(1,0),'se':(1,1),'s':(0,1),'sw':(-1,1),'w':(-1,0),'nw':(-1,-1)}
    adjdict = {'n':(0,-1),'e':(1,0),'s':(0,1),'w':(-1,0)}
    cells = []

    if cell[0] >= len(grid[0])-1:
        adjdict = {k:v for k,v in adjdict.items() if not 'e' in k}
    if cell[0] <= 0:
        adjdict = {k:v for k,v in adjdict.items() if not 'w' in k}
    if cell[1] >= len(grid)-1:
        adjdict = {k:v for k,v in adjdict.items() if not 's' in k}
    if cell[1] <= 0:
        adjdict = {k:v for k,v in adjdict.items() if not 'n' in k}
    
    for neigh in adjdict.values():
        if int(grid[neigh[1]+cell[1]][neigh[0]+cell[0]]) < 9:
            cells.append((neigh[0]+cell[0],neigh[1]+cell[1]))
    return cells 

def astar(x,y,map):
    while len(openlist) > 0:
        closedlist.append((x,y))
        x, y = openlist[0]
        children = get_adjacent_cells(map, (x,y))
        children = [item for item in children if item not in closedlist]
        if len(children) == 0:
            openlist.remove(openlist[0])
        else:
            for c in children:
                if not c in closedlist:
                    if c in openlist:
                        astar(c[0], c[1], map)
                    else:
                        openlist.insert(0, c)
    return len(set(closedlist))

areas = []

for mini in minis:
    openlist = [mini[0]]
    closedlist = []
    amount = astar(mini[0][0],mini[0][1], map)
    areas.append(amount)

areas.sort()
print(f"result part 2: {math.prod(areas[-3:])}")
