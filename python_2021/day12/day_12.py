import copy

with open("day_12_test.txt") as f:
    paths = [x.split('-') for x in f.read().splitlines()]

mapdict = {}
for path in paths:
    if path[0].islower():
        if path[0] in mapdict:
            mapdict[path[0]].append(path[1])
        else:
            mapdict[path[0]] = [path[1]]
    else:
        if path[0] in mapdict:
            mapdict[path[0]].append(path[1])
        else:
            mapdict[path[0]] = [path[1]]
        if path[1] in mapdict:
            mapdict[path[1]].append(path[0])
        else:
            mapdict[path[1]] = [path[0]]

def astar(start,stop, path, mapdict, openlist, closedlist):
    while len(openlist) > 0:
        closedlist.append(start)
        node = openlist[0]
        if node in mapdict:
            children = mapdict[node]
            children = [item for item in children if item not in closedlist]
        else:
            children = []
        if len(children) == 0:
            openlist.remove(openlist[0])
        elif start == stop:
            allpaths.append(copy.deepcopy(path))
        else:
            for c in children:
                if not c in closedlist:
                    if c in openlist:
                        astar(c, stop, path, mapdict, openlist, closedlist)
                    else:
                        openlist.insert(0, c)
    return

openlist = ['start']
closedlist = []
allpaths = []
astar('start','end', path, mapdict, openlist, closedlist)

print(f"result part 2: {math.prod(areas[-3:])}")
