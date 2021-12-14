with open("day_13.txt") as f:
    alltext = f.read()
    dots, folds = alltext.split('\n\n')
    dots = [x.split(',') for x in dots.split('\n')]
    dots = [(int(i[0]), int(i[1])) for i in dots]
    folds = [x.split('=') for x in folds.split('\n')]
    
# dots = [(6,10),(0,14),(9,10),(0,3),(10,4),(4,11),(6,0),(6,12),(4,1),(0,13),(10,12),(3,4),(3,0),(8,4),(1,10),(2,14),(8,10),(9,0)]

def fold(f, pos, grid):
    gridx = max([x[0] for x in grid])+1
    gridy = max([x[1] for x in grid])+1
    if f == 'x':
        for x in range(1, int(((gridx-1)/2)+1)):
            for y in range(0, gridy+1):
                if (pos+x,y) in grid:
                    grid.append((pos-x,y))
                    grid.remove((pos+x,y))
    else:
        for x in range(0, gridx+1):
            for y in range(0, int(((gridy-1)/2)+1)):
                if (x,y+pos) in grid:
                    grid.append((x,y-pos))
                    grid.remove((x,y+pos))
    return grid

# part 1
# newdots = (fold('y',7,dots))
# newdots = (fold('x',655,dots)).copy()
# print(len(set(newdots)))

# part 2

newdots = dots.copy()
for do in folds:
    foldtype = do[0][-1]
    foldpos = int(do[1])
    # print(newdots)
    newdots = (fold(foldtype,foldpos,newdots)).copy()
    print(newdots)

gridx = max([x[0] for x in newdots])+1
gridy = max([x[1] for x in newdots])+1

for y in range(0, gridy):
    line = ''
    for x in range(0, gridx):
        if (x,y) in newdots:
            line = line + "#"
        else:
            line = line + '.'
    print(line)
print("\n")
