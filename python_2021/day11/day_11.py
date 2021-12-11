import numpy as np
import scipy.signal

with open("day_11.txt") as f:
    input = f.read().splitlines()
    octopus = [[int(i) for i in line] for line in input]
    octopus = np.array(octopus)

# octopus = np.array([[1,1,1,1,1],[1,9,9,9,1],[1,9,1,9,1],[1,9,9,9,1],[1,1,1,1,1]])

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
        cells.append(cell + np.array(neigh))
    return cells 


adder = np.full([10,10], 1, dtype='int32')
# adder = np.full([5,5], 1, dtype='int32')

flashes = 0
step = 0
while True:
    step += 1
    octopus = octopus + adder
    flasher = np.argwhere(octopus > 9)
    for i in flasher:
        octopus[i[0]][i[1]] = 0
    # print(octopus)
    # print(flasher)
    while len(flasher) > 0:
        for flash in flasher:
            # octopus[flash[0]][flash[1]] = 0
            adjcells = get_adjacent_cells(octopus, flash)
            for cell in adjcells:
                if not octopus[cell[0]][cell[1]] == 0:
                    octopus[cell[0]][cell[1]] += 1
            flashes += 1
        flasher = np.argwhere(octopus > 9)
        for i in flasher:
            octopus[i[0]][i[1]] = 0
    # print(octopus)
    if step == 100:
        print(f"result part 1: {flashes}")
    if np.sum(octopus) == 0:
        print(f"result part 2: {step}")
        break   



