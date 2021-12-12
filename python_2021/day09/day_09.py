import re
import numpy as np

with open("day_09.txt") as f:
    input = f.read()
    map = input.splitlines()
    maplist = [int(i) for i in input.replace('\n','')]

#dimy, dimx = len(map), int(math.log10(map[0]))+1
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

# map = [re.sub('[0-8]', '0', i) for i in map]

def get_adjacent_cells(grid, cell):
    adjdict = {'n':(0,-1),'ne':(1,-1),'e':(1,0),'se':(1,1),'s':(0,1),'sw':(-1,1),'w':(-1,0),'nw':(-1,-1)}
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
        if int(grid[neigh[0]][neigh[1]]) < 9:
            print(grid[neigh[0]][neigh[1]])
            cells.append(neigh)
    return cells 

def count(x, y, map):
    counter = 0
    children = get_adjacent_cells(map, (x,y))
    children = [i for i in children if int(map[i[1]][i[0]]) < 9]
    if len(children) > 0:
        for c in children:
            counter += count(x+c[0],y+c[1], map)
    else:
        counter += 1
    return counter

areas = []

for mini in minis:
    areas.append(count(mini[0][1],mini[0][0], map))



print(f"result part 2: {sum([i[1]+1 for i in minis])}")
