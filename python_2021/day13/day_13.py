with open("day_13.txt") as f:
    alltext = f.read()
    dots, folds = alltext.split('\n\n')
    dots = [x.split(',') for x in dots.split('\n')]
    dots = [[int(i[0]), int(i[1])] for i in dots]
    folds = [x.split('=') for x in folds.split('\n')]
    
dots = [(6,10),(0,14),(9,10),(0,3),(10,4),(4,11),(6,0),(6,12),(4,1),(0,13),(10,12),(3,4),(3,0),(8,4),(1,10),(2,14),(8,10),(9,0)]

def fold(f, pos, grid):
    gridx = max([x[0] for x in dots])
    gridy = max([x[1] for x in dots])
    if f == 'x':
        for x in range(1, int(((gridx-1)/2)+1)):
            for y in range(0, gridy):
                print(pos+x,y,[pos-x,y])
                print(grid)
                if (pos+x,y) in grid:
                    grid.append([pos-x,y])
                    grid.remove((pos+x,y))
    else:
        pass
    return grid

#newgrid = (fold('x',655,grid))
newdots = (fold('x',7,dots))

print(len(newdots))