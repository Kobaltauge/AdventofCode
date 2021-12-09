#import math

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
    if y == dimy:
        del dir['down']
    if x == 0:
        del dir['left']
    if x == dimx:
        del dir['right']
    return [map[(x + dir[i][1])][(y + dir[i][0])] for i in dir]

for y in range(0, dimy):
    for x in range(0, dimx):
        neigh = get_neigh(x,y, map)
        min = min(neigh)
        print(map[x][y], min)
        print("x")

# for i in range(0, max):
#     if i > 0:
#         neigh.append(maplist[i-1])
#     if i < max:
#     if ((i - dimx) < 0) and ((i + dimx) > max):
