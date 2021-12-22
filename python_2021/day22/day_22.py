from  itertools import chain

with open("day_22.txt") as f:
    actions = [[j[0],j[1].split(",")] for j in [i.replace('\n','').split(" ") for i in f.readlines()]]

class cube:
    def __init__(self,x,y,z):
        self.x = x
        self.y = y
        self.z = z
        self.state = "off"
    def __iter__(self):
        yield 'x', self.x
        yield 'y', self.y
        yield 'z', self.z
        yield 'state', self.state

cubes = {}
for action in actions:
    kdict = {}
    for k in action[1]:
        i = k.split('=')
        kdict[i[0]] = list(map(int, i[1].split("..")))
    if not ((min(chain(*kdict.values())) < -50) or (max(chain(*kdict.values())) > 50)):
        for xi in range(kdict['x'][0], kdict['x'][1]+1):
            for yi in range(kdict['y'][0], kdict['y'][1]+1):
                for zi in range(kdict['z'][0], kdict['z'][1]+1):
                    name = str(xi)+","+str(yi)+","+str(zi)
                    if not name in cubes:
                        cubes[name] = cube(xi,yi,zi)
                    if not cubes[name].state == action[0]:
                        cubes[name].state = action[0]

print(len([i for i in cubes.values() if i.state == "on"]))
print("\n")

cubes = {}
for action in actions:
    print(len(cubes))
    kdict = {}
    for k in action[1]:
        i = k.split('=')
        kdict[i[0]] = list(map(int, i[1].split("..")))
    for xi in range(kdict['x'][0], kdict['x'][1]+1):
        for yi in range(kdict['y'][0], kdict['y'][1]+1):
            for zi in range(kdict['z'][0], kdict['z'][1]+1):
                name = str(xi)+","+str(yi)+","+str(zi)
                if not name in cubes:
                    cubes[name] = cube(xi,yi,zi)
                if not cubes[name].state == action[0]:
                    cubes[name].state = action[0]

print(len([i for i in cubes.values() if i.state == "on"]))
print("\n")