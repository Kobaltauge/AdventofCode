#import math
import re

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

map = [re.sub('[0-8]', '0', i) for i in map]

print(f"result part 2: {sum([i[1]+1 for i in minis])}")
