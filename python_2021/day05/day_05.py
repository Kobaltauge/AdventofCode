with open("day_05.txt") as f:
    # input = list(f.readline().split(' -> '))
    koords = [list(map(int, i.replace(' -> ',',').replace('\n','').split(','))) for i in f.readlines()]

def linecoords(x1,y1,x2,y2):
    if (x1 == x2) or (y1 == y2):
        xs = 1 if x1<x2 else -1
        ys = 1 if y1<y2 else -1
        values = []
        for i in range(x1, x2+xs, xs):
            for j in range(y1, y2+ys, ys):
                values.append([i,j])
    else:
        values = []
    return values

max = sorted([item for sublist in koords for item in sublist])[-1] + 1
pipemap = [[0]*max for _ in range(max)]

for koord in koords:
    line = linecoords(koord[0],koord[1],koord[2],koord[3])
    for point in line:
        pipemap[point[0]][point[1]] += 1
print(f"result part 1: {len([item for sublist in pipemap for item in sublist if item >= 2])}")

## Part 2

def linecoords(x1,y1,x2,y2):
    if (x1 == x2) or (y1 == y2):
        xs = 1 if x1<x2 else -1
        ys = 1 if y1<y2 else -1
        values = []
        for i in range(x1, x2+xs, xs):
            for j in range(y1, y2+ys, ys):
                values.append([i,j])
    else:
        values = []
        xs = 1 if x1<x2 else -1
        ys = 1 if y1<y2 else -1
        xes = [i for i in range(x1, x2+xs, xs)]
        yes = [i for i in range(y1, y2+ys, ys)]
        for i in range(0, len(xes)):
            values.append([xes[i], yes[i]])
    return values

max = sorted([item for sublist in koords for item in sublist])[-1] + 1
pipemap = [[0]*max for _ in range(max)]

for koord in koords:
    line = linecoords(koord[0],koord[1],koord[2],koord[3])
    for point in line:
        pipemap[point[0]][point[1]] += 1
print(f"result part 2: {len([item for sublist in pipemap for item in sublist if item >= 2])}")
